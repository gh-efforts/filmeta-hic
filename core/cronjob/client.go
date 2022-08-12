package cronjob

func SetupCron(opts ...Option) *Runner {
	cron := NewRunner(opts...)
	return cron
}
