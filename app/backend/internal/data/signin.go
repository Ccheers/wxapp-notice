package data

import (
	"context"
	"fmt"
	"time"

	"github.com/Ccheers/wxapp-notice/app/backend/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"go.etcd.io/bbolt"
)

type SignInRepoImpl struct {
	data *Data
	log  *log.Helper
}

var _ biz.SignInRepo = (*SignInRepoImpl)(nil)

func (s *SignInRepoImpl) TableName() string {
	return "signin"
}

func (s *SignInRepoImpl) TableNameWithOpenID(openid string) string {
	return fmt.Sprintf("%s:%s", s.TableName(), openid)
}

func NewSignInRepoImpl(data *Data, logger log.Logger) biz.SignInRepo {
	return &SignInRepoImpl{data: data, log: log.NewHelper(logger)}
}

func (s *SignInRepoImpl) SignIn(ctx context.Context, openid string, date time.Time) error {
	return s.data.db.Update(func(tx *bbolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(s.TableNameWithOpenID(openid)))
		if err != nil {
			return err
		}
		return b.Put([]byte(date.Format(biz.SignInTimeFormat)), []byte("1"))
	})
}

func (s *SignInRepoImpl) GetSignInList(ctx context.Context, openid string) ([]time.Time, error) {
	var list []time.Time
	err := s.data.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(s.TableNameWithOpenID(openid)))
		if b == nil {
			return errBucketNotExist
		}
		return b.ForEach(func(k, v []byte) error {
			t, err := time.Parse(biz.SignInTimeFormat, string(k))
			if err != nil {
				return err
			}
			list = append(list, t)
			return nil
		})
	})
	return list, err
}
