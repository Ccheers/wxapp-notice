package biz

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	pb "github.com/Ccheers/wxapp-notice/app/backend/api/notice/v1"
	v1 "github.com/Ccheers/wxapp-notice/app/backend/api/wx/v1"
	"github.com/Ccheers/wxapp-notice/app/backend/internal/conf"
	"github.com/Ccheers/wxapp-notice/app/backend/internal/cron"
	"github.com/go-kratos/kratos/v2/log"
	cron2 "github.com/robfig/cron/v3"
)

const subTmplID = "TPxDR87VAsidwvX6UKdlVPbKVdwPl4qnhxRV8qFoAYA"

type JobFuncID uint

const (
	Once     = iota // 一次
	WorkDay         // 工作日
	Week            // 每周
	Month           // 每月
	EveryDay        // 每天
)

type JobType int

const (
	JobTypeOnce = -1
	JobTypeCron = 1
)

const JobFuncNotice JobFuncID = 1

type CallbackFunc func(job *Job) error

var (
	errRegisterJobFunc = errors.New("register job func error")
	errGetJobFunc      = errors.New("get job func error")
	ErrNoticeNotFound  = errors.New("notice not found")
	ErrGetAccessToken  = errors.New("get access token error")
	ErrSendNotice      = errors.New("send notice error")
	errFixNotice       = errors.New("fix notice error")
)

func IsErrNoticeNotFound(err error) bool {
	return errors.Is(err, ErrNoticeNotFound)
}

type JobRepo interface {
	PutJob(ctx context.Context, job *Job) (*Job, error)
	DeleteJob(ctx context.Context, openid string, jobID uint64) error
	GetAllJobs(ctx context.Context, openid string) ([]*Job, error)
	GetJobByID(ctx context.Context, openid string, jobID uint64) (*Job, error)
}

type Job struct {
	ID              uint64 `json:"id"`
	Openid          string `json:"from_user_id"`
	Content         string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Time            string `protobuf:"bytes,4,opt,name=time,proto3" json:"time,omitempty"`
	Frequency       int64  `protobuf:"bytes,5,opt,name=frequency,proto3" json:"frequency,omitempty"`
	FrequencyDetail string `protobuf:"bytes,6,opt,name=frequency_detail,json=frequencyDetail,proto3" json:"frequency_detail,omitempty"`

	CronID      cron2.EntryID `json:"cron_id"`
	CronExpress string        `json:"cron_express"`

	JobType   JobType   `json:"job_type"`
	JobFuncID JobFuncID `json:"job_func_id"`
}

func (j *Job) Fixed() error {

	switch j.Frequency {
	case Once:
		j.JobType = JobTypeOnce

		if j.FrequencyDetail == "0" {
			j.CronExpress = "* * *"
		} else {
			t, err := time.ParseInLocation("2006-01-02", j.FrequencyDetail, time.Local)
			if err != nil {
				return fmt.Errorf("%w: %s", errFixNotice, err)
			}
			j.CronExpress = fmt.Sprintf("%s * *", t.Format("2"))
		}

	case WorkDay:
		j.JobType = JobTypeCron
		j.CronExpress = "* * 1,2,3,4,5"
	case Week:
		j.JobType = JobTypeCron
		j.CronExpress = fmt.Sprintf("* * %s", j.FrequencyDetail)
	case Month:
		j.JobType = JobTypeCron
		j.CronExpress = fmt.Sprintf("%s * *", j.FrequencyDetail)
	case EveryDay:
		j.JobType = JobTypeCron
		j.CronExpress = "* * *"
	default:
		return fmt.Errorf("%w: %s", errFixNotice, "未知的 job 类型")
	}

	t, err := time.ParseInLocation("15:04", j.Time, time.Local)
	if err != nil {
		return fmt.Errorf("%w: %s", errFixNotice, err)
	}
	j.CronExpress = fmt.Sprintf("%s %s", t.Format("4 15"), j.CronExpress)

	j.JobFuncID = JobFuncNotice

	return nil
}

type NoticeRepo interface {
	AddOpenID(ctx context.Context, openID string) error
	CheckOpenID(ctx context.Context, openID string) error
}

type NoticeUseCase struct {
	jobRepo         JobRepo
	noticeRepo      NoticeRepo
	wxRepo          WxRepo
	log             *log.Helper
	callbackFuncMap sync.Map
	cron            *cron.Cron
}

func NewNoticeUseCase(jobRepo JobRepo, wxRepo WxRepo, noticeRepo NoticeRepo, logger log.Logger, c *cron.Cron, wxConf *conf.WxAppConfig) *NoticeUseCase {
	n := &NoticeUseCase{
		jobRepo:         jobRepo,
		noticeRepo:      noticeRepo,
		log:             log.NewHelper(logger),
		callbackFuncMap: sync.Map{},
		cron:            c,
		wxRepo:          wxRepo,
	}
	err := n.RegisterJobFunc(JobFuncNotice, callback(wxRepo, log.NewHelper(logger), wxConf))
	if err != nil {
		panic(err)
	}
	return n
}

func (n *NoticeUseCase) SubscribeNotice(ctx context.Context, request *pb.SubscribeNoticeRequest) (err error) {
	return n.noticeRepo.AddOpenID(ctx, request.GetOpenid())
}

func (n *NoticeUseCase) CheckOpenID(ctx context.Context, request *pb.CheckSubscribeStatusRequest) (b bool, err error) {
	err = n.noticeRepo.CheckOpenID(ctx, request.GetOpenid())
	if err != nil {
		if IsErrNoticeNotFound(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (n *NoticeUseCase) AddNotice(ctx context.Context, job *Job) (*Job, error) {
	err := job.Fixed()
	if err != nil {
		return nil, err
	}

	job, err = n.jobRepo.PutJob(ctx, job)
	if err != nil {
		return nil, err
	}
	entryID, err := n.cron.AddCron(job.CronExpress, n.WithCronFunc(job.Openid, job.ID))
	if err != nil {
		return nil, err
	}

	job.CronID = entryID
	if err != nil {
		return nil, err
	}

	return n.jobRepo.PutJob(ctx, job)
}

func (n *NoticeUseCase) DeleteJob(ctx context.Context, openID string, jobID uint64) error {
	return n.jobRepo.DeleteJob(ctx, openID, jobID)
}

func (n *NoticeUseCase) WithCronFunc(openid string, jobID uint64) func() {
	return func() {
		n.log.Debugf("job %d start", jobID)
		job, err := n.jobRepo.GetJobByID(context.TODO(), openid, jobID)
		if err != nil {
			n.log.Errorf("get job by id error: %v", err)
			return
		}

		// 如果是一次性任务，则在任务列表中删除
		if job.JobType == JobTypeOnce {

			err := n.DeleteJob(context.TODO(), job.Openid, jobID)
			if err != nil {
				n.log.Error(err.Error())
			}

			n.cron.DelCron(job.CronID)
		}

		f, err := n.GetJobFunc(job.JobFuncID)
		if err != nil {
			n.log.Error(err.Error())
			return
		}
		err = f(job)
		if err != nil {
			n.log.Error(err.Error())
			return
		}
	}
}

func (n *NoticeUseCase) RegisterJobFunc(typeID JobFuncID, f CallbackFunc) error {
	_, ok := n.callbackFuncMap.Load(typeID)
	if ok {
		return fmt.Errorf("%w: job type %d already registered", errRegisterJobFunc, typeID)
	}
	n.callbackFuncMap.Store(typeID, f)
	return nil
}

func (n *NoticeUseCase) GetJobFunc(typeID JobFuncID) (CallbackFunc, error) {
	jFunc, ok := n.callbackFuncMap.Load(typeID)
	if !ok {
		return nil, fmt.Errorf("%w: job type %d not registered", errGetJobFunc, typeID)
	}
	return jFunc.(CallbackFunc), nil
}

// 注册到 CRON 的 job，周期性任务到期会回调这个函数
func callback(repo WxRepo, log *log.Helper, wxConf *conf.WxAppConfig) CallbackFunc {
	return func(job *Job) error {
		reply, err := repo.GetAccessToken(context.TODO(), &v1.GetAccessTokenRequest{
			GrantType: "client_credential",
			Appid:     wxConf.AppId,
			Secret:    wxConf.AppSecret,
		})
		if err != nil {
			return err
		}
		if reply.Errcode != 0 {
			return fmt.Errorf("%w: %s", ErrGetAccessToken, reply.Errmsg)
		}

		res, err := repo.SubscribeSend(context.TODO(), &v1.SubscribeSendRequest{
			AccessToken: reply.AccessToken,
			Touser:      job.Openid,
			TemplateId:  subTmplID,
			Data: map[string]*v1.TmplData{
				"thing1": {
					Value: job.Content,
				},
				"time3": {
					Value: time.Now().Format("2006.01.02 15:04"),
				},
			},
		})
		if err != nil {
			return err
		}
		if reply.Errcode != 0 {
			return fmt.Errorf("%w: %s", ErrSendNotice, reply.Errmsg)
		}
		log.Infof("SubscribeSend result %v", res)
		return nil
	}
}
