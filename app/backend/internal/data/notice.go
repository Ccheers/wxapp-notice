package data

import (
	"context"

	"github.com/Ccheers/wxapp-notice/app/backend/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"go.etcd.io/bbolt"
)

type NoticeRepoImpl struct {
	data *Data
	log  *log.Helper
}

func NewNoticeRepoImpl(data *Data, logger log.Logger) biz.NoticeRepo {
	return &NoticeRepoImpl{data: data, log: log.NewHelper(logger)}
}

func (n *NoticeRepoImpl) TableName() string {
	return "notice"
}

func (n *NoticeRepoImpl) AddOpenID(ctx context.Context, openID string) error {
	return n.data.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(n.TableName()))

		err := b.Put([]byte(openID), []byte(openID))
		if err != nil {
			return err
		}
		return nil
	})
}

func (n *NoticeRepoImpl) CheckOpenID(ctx context.Context, openID string) error {
	return n.data.db.View(func(tx *bbolt.Tx) error {
		n.log.Infof("openID: %s", openID)

		b := tx.Bucket([]byte(n.TableName()))
		res := b.Get([]byte(openID))
		if len(res) == 0 {
			return biz.ErrNoticeNotFound
		}
		return nil
	})
}

func (n *NoticeRepoImpl) GetAllOpenID(ctx context.Context) ([]string, error) {
	ids := make([]string, 0)
	err := n.data.db.View(func(tx *bbolt.Tx) error {
		c := tx.Bucket([]byte(n.TableName())).Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			ids = append(ids, string(v))
		}
		return nil
	})
	return ids, err
}
