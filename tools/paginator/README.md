### 统一分页器


#### demo
```go
package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/bitrainforest/filmeta-hic/tools/paginator"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/bitrainforest/filmeta-hic/store/mongox"
)

type (
	List struct {
		userList []User
		paginator.ListPage
	}

	User struct {
		UserId   string `json:"user_id" bson:"user_id"`
		UserName string `json:"user_name" bson:"user_name"`
		Email    string `json:"email" bson:"email"`
	}
)

func main() {
	c := mongox.Conf{Uri: "mongodb://127.0.0.1:27017/"}
	client, err := c.GetClient()
	if err != nil || client==nil {
		// just for demo 
		panic(err)
	}
	
	
	p := paginator.Page{
		Page:     2, 
		PageSize: 1,
	}
	
	var (
		total int64
	)
	filter := bson.M{"user_id": "test-01"}
	
	collection:= func()*mongo.Collection {
		return client.Database("filmeta").Collection("user")
	}

	var (
		cur *mongo.Cursor
	)
	if cur, err = collection().Find(context.TODO(), filter, p.Paginate()); err != nil {
		panic(err)
	}

	var (
		list List
	)
	if err = cur.All(context.TODO(), &list.userList); err != nil {
		panic(err)
	}

	// total
	if total, err = collection().
		CountDocuments(context.Background(), filter); err != nil {
		// just for demo
		panic(err)
	}
	list.ListPage = p.ToListPage(total)
	fmt.Printf("%+v", list)
}
```