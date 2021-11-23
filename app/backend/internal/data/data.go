package data

import (
	"fmt"
	"os"

	"github.com/Ccheers/wxapp-notice/app/backend/internal/conf"
	"go.etcd.io/bbolt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewWxRepoImpl, NewJobRepoImpl, NewNoticeRepoImpl)

type BucketName interface {
	TableName() string
}

func registerBucketNames() []BucketName {
	return []BucketName{
		new(JobRepoImpl),
	}
}

// Data .
type Data struct {
	db  *bbolt.DB
	log *log.Helper
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	err := os.MkdirAll(c.Database.DbPath, os.ModePerm)
	if err != nil {
		panic(err)
	}
	db, err := bbolt.Open(fmt.Sprintf("%s/%s", c.Database.DbPath, "my.db"), 0600, nil)
	if err != nil {
		panic(err)
	}

	helper := log.NewHelper(logger)
	for _, name := range registerBucketNames() {
		err = db.Update(func(tx *bbolt.Tx) error {
			_, err := tx.CreateBucketIfNotExists([]byte(name.TableName()))
			if err != nil {
				helper.Errorf("create bucket error: %v", err)
			}
			return nil
		})
		if err != nil {
			helper.Errorf("create bucket error: %v", err)
		}
	}
	cleanup := func() {
		helper.Info("closing the data resources")
	}
	return &Data{
		db:  db,
		log: helper,
	}, cleanup, nil
}
