package data

import (
	"context"
	"encoding/json"

	"github.com/Ccheers/wxapp-notice/app/backend/internal/biz"
	"github.com/Ccheers/wxapp-notice/app/backend/internal/pkg/itob"
	"github.com/go-kratos/kratos/v2/log"
	"go.etcd.io/bbolt"
)

type JobRepoImpl struct {
	data *Data
	log  *log.Helper
}

func NewJobRepoImpl(data *Data, logger log.Logger) biz.JobRepo {
	return &JobRepoImpl{data: data, log: log.NewHelper(logger)}
}

func (j *JobRepoImpl) PutJob(ctx context.Context, job *biz.Job) (*biz.Job, error) {
	err := j.data.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(j.TableName()))

		if job.ID == 0 {
			id, err := b.NextSequence()
			if err != nil {
				return err
			}
			job.ID = id
		}

		bts, err := json.Marshal(job)
		if err != nil {
			return err
		}
		err = b.Put(itob.Itob(job.ID), bts)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return job, nil
}

func (j *JobRepoImpl) DeleteJob(ctx context.Context, jobID uint64) error {
	return j.data.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(j.TableName()))
		return b.Delete(itob.Itob(jobID))
	})
}

func (j *JobRepoImpl) GetAllJobs(ctx context.Context) ([]*biz.Job, error) {
	jobs := make([]*biz.Job, 0)
	err := j.data.db.View(func(tx *bbolt.Tx) error {
		c := tx.Bucket([]byte(j.TableName())).Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			job := new(biz.Job)
			err := json.Unmarshal(v, job)
			if err != nil {
				j.log.Errorf(err.Error())
				continue
			}
			jobs = append(jobs, job)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return jobs, nil
}

func (j *JobRepoImpl) GetJobByID(ctx context.Context, jobID uint64) (job *biz.Job, err error) {
	job = new(biz.Job)

	err = j.data.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(j.TableName()))

		v := b.Get(itob.Itob(jobID))
		j.log.Infof("v: %s", string(v))
		err := json.Unmarshal(v, job)
		if err != nil {
			return err
		}
		return nil
	})
	return
}

func (j *JobRepoImpl) TableName() string {
	return "job"
}
