package mongox

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	defaultTimeout = time.Second * 5
	MinPool        = 80
	DefaultMaxPool = 400
	AllowMaxPool   = 600
)

type (
	Conf struct {
		Uri      string `json:"uri"`      // example: mongodb://localhost:27017
		WorkPool uint64 `json:"workPool"` // mongodb connection pool size
		Register *bsoncodec.Registry
	}
)

func (conf *Conf) GetClient() (*mongo.Client, error) {
	if conf.Uri == "" {
		return nil, fmt.Errorf("empty mongo uri")
	}
	opts := new(options.ClientOptions)
	opts.SetConnectTimeout(defaultTimeout)
	opts.SetServerSelectionTimeout(defaultTimeout)
	opts.SetMinPoolSize(MinPool)
	opts.SetMaxPoolSize(DefaultMaxPool)
	if conf.WorkPool > MinPool && conf.WorkPool <= AllowMaxPool {
		opts.SetMaxPoolSize(conf.WorkPool)
	}

	opts.ApplyURI(conf.Uri)
	var (
		err    error
		client *mongo.Client
	)
	if conf.Register != nil {
		opts.SetRegistry(conf.Register)
	}
	client, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, fmt.Errorf("mongo connect err: %v", err)
	}

	// ping client conn
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*2)
	defer cancel()
	if err = client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("ping mongo err:%v", err)
	}
	return client, nil
}
