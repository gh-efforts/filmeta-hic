package helper

import (
	"fmt"
	"github.com/bitrainforest/filmeta-hic/core/log"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	MongoDBErr = fmt.Errorf("internal database errors")
)

func WarpMongoErr(err error) error {
	if err != nil && errors.Is(err, mongo.ErrNoDocuments) {
		return nil
	}
	if log.IsInit() {
		log.Errorf("[WarpMongoErr] err: %v", err)
	}
	return MongoDBErr
}
