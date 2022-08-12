package cronjob

import (
	"context"
	"fmt"
	"testing"
)

func fn1(ctx context.Context) {
	fmt.Println("fn1")
}
func fn2(ctx context.Context) {
	fmt.Println("fn2")
}

func TestCronJob_CronJobRun(t *testing.T) {
	job := NewRunner(BlockOption(func() bool {
		// todo
		return true
	}))

	job.AddJob(JobInfo{
		Spec: "@every 5s",
		Name: "fn1",
		Fn:   fn1,
	})

	job.AddJob(JobInfo{
		Spec: "* * * * * ?",
		Name: "fn2",
		Fn:   fn2,
	})

	err := job.Start(context.TODO())
	if err != nil {
		fmt.Printf(" error: %s \n", err.Error())
		return
	}
}
