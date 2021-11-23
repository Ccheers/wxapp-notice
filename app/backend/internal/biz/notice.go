package biz

import (
	"context"
	"errors"

	pb "github.com/Ccheers/wxapp-notice/app/backend/api/notice/v1"
	"github.com/robfig/cron/v3"
)

type FreType string

type JobFuncID uint

const (
	None    FreType = "* * *"
	ADay    FreType = "* * *"
	AWeek   FreType = "* * %d"
	AMonth  FreType = "%d * *"
	WorkDay FreType = "* * 1,2,3,4,5"
)

type JobType int

const (
	JobTypeOnce = -1
	JobTypeCron = 1
)

type CallbackFunc func(job *Job) error

var (
	errRegisterJobFunc = errors.New("register job func error")
	errGetJobFunc      = errors.New("get job func error")
	ErrNoticeNotFound  = errors.New("notice not found")
)

func IsErrNoticeNotFound(err error) bool {
	return errors.Is(err, ErrNoticeNotFound)
}

type JobRepo interface {
	PutJob(ctx context.Context, job *Job) (*Job, error)
	DeleteJob(ctx context.Context, jobID uint64) error
	GetAllJobs(ctx context.Context) ([]*Job, error)
	GetJobByID(ctx context.Context, jobID uint64) (*Job, error)
}

type Job struct {
	ID          uint64       `json:"id"`
	OpenID      string       `json:"from_user_id"`
	CronID      cron.EntryID `json:"cron_id"`
	CronExpress string       `json:"cron_express"`
	Content     string       `json:"content"`
	JobType     JobType      `json:"job_type"`
	JobFuncID   JobFuncID    `json:"job_func_id"`
}

type NoticeRepo interface {
	AddOpenID(ctx context.Context, openID string) error
	CheckOpenID(ctx context.Context, openID string) error
}

type NoticeUseCase struct {
	jobRepo    JobRepo
	noticeRepo NoticeRepo
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
