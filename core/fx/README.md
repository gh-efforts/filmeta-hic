CONY FROM go-zero !!!!!!

```go
package main

import (
	"fmt"
	"github.com/bitrainforest/filmeta-hic/core/fx"
	"time"
)

func main() {
	n := time.Now()
	userList := getUser()
	fx.From(func(source chan<- interface{}) {
		for _, user := range userList {
			source <- user
		}
	}).Filter(func(item interface{}) bool {
		user := item.(User)
		if user.Age < 10 {
			return false
		}
		return true
	}).Walk(func(item interface{}, pipe chan<- interface{}) {
		user := item.(User)
		user.Money = 200
		if user.Age == 35 {
			user.Money = -100
		}
		pipe <- user
	}).ForEach(func(item interface{}) {
		fmt.Printf("user:%v\n", item)
	})
	fmt.Printf("time spent:%v\n", time.Since(n).Seconds())

}

/********test user **********/

type User struct {
	Name  string
	Age   uint8
	Money float64
}

func getUser() []User {
	return []User{
		{Age: 5, Name: "curry"},
		{Age: 10, Name: "remember"},
		{Age: 18, Name: "test1"},
		{Age: 35, Name: "ai"},
	}
}

```

针对 fx,还有以下这些操作:

API 作用

- Distinct(fn)    fn中选定特定item类型，对其去重
- Filter(fn, option)    fn指定具体规则，满足规则的element传递给下一个 stream
- Group(fn)    根据fn把stream中的element分到不同的组中
- Head(num)    取出stream中前 num 个element ，生成一个新的stream
- Map(fn, option)    将每个ele转换为另一个对应的ele， 传递给下一个 stream
- Merge()    将所有ele合并到一个slice中并生成一个新stream
- Reverse()    反转stream中的element。【使用双指针】
- Sort(fn)    按照 fn 排序stream中的element
- Tail(num)    取出stream最后的 num 个element，生成一个新 stream。【使用双向环状链表】
- Walk(fn, option)    把 fn 作用在 source 的每个元素。生成新的 stream

不再生成新的 stream，做最后的求值操作：

- ForAll(fn)    按照fn处理stream，且不再产生stream【求值操作】
- ForEach(fn)    对 stream 中所有 element 执行fn【求值操作】
- Parallel(fn, option)    将给定的fn与给定的worker数量并发应用于每个element【求值操作】
- Reduce(fn)    直接处理stream【求值操作】
- Done()    啥也不做，等待所有的操作完成