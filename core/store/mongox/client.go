package mongox

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	defaultTimeout = time.Second * 5
)

type (
	Conf struct {
		Uri string `json:"uri"` // example: mongodb://localhost:27017
	}
)

func (conf *Conf) GetClient() (*mongo.Client, error) {
	if conf.Uri == "" {
		return nil, fmt.Errorf("empty mongo uri")
	}
	opts := new(options.ClientOptions)
	opts.SetConnectTimeout(defaultTimeout)
	opts.SetServerSelectionTimeout(defaultTimeout)
	opts.ApplyURI(conf.Uri)
	var (
		err    error
		client *mongo.Client
	)
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
