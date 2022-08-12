package cronjob

import (
	"context"
	"github.com/bitrainforest/filmeta-hic/core/log"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/robfig/cron/v3"
	"time"
)

type (
	JobInfo struct {
		Name string
		Spec string
		Fn   JobFn
	}
	JobFn func(ctx context.Context)

	Runner struct {
		blocked bool
		Jobs    []*JobInfo

		cron *cron.Cron
	}

	Option func(job *Runner)
)

var _ transport.Server = (*Runner)(nil)

func BlockOption(blockFn func() bool) Option {
	return func(job *Runner) {
		if blockFn() {
			job.blocked = true
		}
	}
}

func NewRunner(options ...Option) *Runner {
	c := &Runner{
		cron: cron.New(cron.WithParser(cron.NewParser(
			cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor))),
	}

	for _, v := range options {
		v(c)
	}

	return c
}

func (c *Runner) AddJob(job JobInfo) {
	c.Jobs = append(c.Jobs, &job)
}

func warpJob(name string, fn JobFn) cron.Job {
	return cron.NewChain(cron.SkipIfStillRunning(newLogger())).Then(cron.FuncJob(func() {
		ctx := context.Background()
		defer func() {
			if p := recover(); p != nil {
				log.Errorf("定时任务 %s 发生异常: %v", name, p)
			}
		}()

		begin := time.Now()
		log.Errorf("开始执行脚本", name)

		fn(ctx)
		log.Infof("结束执行脚本: %s, 耗时%.3fs", name, time.Since(begin).Seconds())
	}))
}

func (c *Runner) Start(ctx context.Context) (err error) {
	if c.blocked {
		return nil
	}

	for _, job := range c.Jobs {
		_, err = c.cron.AddJob(job.Spec, warpJob(job.Name, job.Fn))
		if err != nil {
			return err
		}
		log.Infof("添加定时任务: %s,周期", job.Name, job.Spec)
	}

	c.cron.Start()

	return
}

func (c *Runner) Stop(ctx context.Context) error {
	c.cron.Stop()
	log.Infof("停止定时任务")
	return nil
}