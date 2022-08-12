## 秒级计划任务

### 例子
```go 
job := NewCronJob()

// 增加任务
job.AddJob(JobInfo{
    Spec: "@every 5s",
    Name: "fn1",
    Fn:   fn1,
})

// 增加任务
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
```
