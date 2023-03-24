package mongodbr

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FindOption func(*options.FindOptions)

func FindOptionWithSkip(skip int64) FindOption {
	return func(fo *options.FindOptions) {
		fo.SetSkip(skip)
	}
}

func FindOptionWithLimit(limit int64) FindOption {
	return func(fo *options.FindOptions) {
		fo.SetLimit(limit)
	}
}

func FindOptionWithSort(sort bson.D) FindOption {
	return func(fo *options.FindOptions) {
		fo.SetSort(sort)
	}
}