package cron

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

//
//
//
//        ***************************     ***************************         *********      ************************
//      *****************************    ******************************      *********      *************************
//     *****************************     *******************************     *********     *************************
//    *********                         *********                *******    *********     *********
//    ********                          *********               ********    *********     ********
//   ********     ******************   *********  *********************    *********     *********
//   ********     *****************    *********  ********************     *********     ********
//  ********      ****************    *********     ****************      *********     *********
//  ********                          *********      ********             *********     ********
// *********                         *********         ******            *********     *********
// ******************************    *********          *******          *********     *************************
//  ****************************    *********            *******        *********      *************************
//    **************************    *********              ******       *********         *********************
//
//

var Provider = wire.NewSet(NewCron)

type Cron struct {
	cron *cron.Cron
	log  *log.Helper
}

func NewCron(logger log.Logger) *Cron {
	c := cron.New(
		cron.WithLogger(cron.VerbosePrintfLogger(logrus.StandardLogger())),
	)
	ins := &Cron{
		cron: c,
		log:  log.NewHelper(logger),
	}
	go ins.Run()
	return ins
}

func (c *Cron) Run() {
	c.log.Info("---------------------CRON START---------------------")
	c.cron.Run()
}

func (c *Cron) AddCron(exp string, job func()) (cron.EntryID, error) {
	id, err := c.cron.AddFunc(exp, job)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (c *Cron) DelCron(id cron.EntryID) {
	c.cron.Remove(id)
}

func (c *Cron) Stop() context.Context {
	return c.cron.Stop()
}

func (c *Cron) Start() {
	c.cron.Start()
}
