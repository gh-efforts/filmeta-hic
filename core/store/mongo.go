package store

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/bitrainforest/filmeta-hic/core/store/mongox"

	"github.com/bitrainforest/filmeta-hic/core/assert"
	"github.com/go-kratos/kratos/v2/config"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	bundle *mongo.Client
)

func GetMongoDB(database string) *mongo.Database {
	return bundle.Database(database)
}

func MustLoadMongoDB(conf config.Config, fn func(cfg config.Config) (*mongox.Conf, error)) {
	if fn == nil {
		assert.CheckErr(errors.New("fn must not nil"))
	}
	mongoCfg, err := fn(conf)
	if err != nil {
		assert.CheckErr(err)
	}
	client, err := mongoCfg.GetClient()
	if err != nil {
		assert.CheckErr(err)
	}
	if client == nil {
		assert.CheckErr(fmt.Errorf("monggo client is nil"))
	}
	bundle = client
}
