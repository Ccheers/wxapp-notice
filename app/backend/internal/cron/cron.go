package cron

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/robfig/cron/v3"
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

type SelfLog struct {
	log log.Logger
}

func NewSelfLog(log log.Logger) *SelfLog {
	return &SelfLog{log: log}
}

func (l *SelfLog) Printf(format string, args ...interface{}) {
	res := append([]interface{}{}, format)
	res = append(res, args...)
	l.log.Log(log.LevelInfo, res...)
}

func NewCron(logger log.Logger) *Cron {
	c := cron.New(
		cron.WithLogger(cron.VerbosePrintfLogger(NewSelfLog(logger))),
	)
	ins := &Cron{
		cron: c,
		log:  log.NewHelper(logger),
	}
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

func (c *Cron) Start(ctx context.Context) error {
	c.Run()
	return nil
}

func (c *Cron) Stop(ctx context.Context) error {
	ctx = c.cron.Stop()
	// wait job done
	select {
	case <-ctx.Done():
	}
	return nil
}
