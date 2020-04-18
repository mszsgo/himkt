package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CollectionInterface interface {
	Collection() *mongo.Collection
}

func Find(c CollectionInterface, result interface{}, ctx context.Context, filter interface{}, opts ...*options.FindOptions) (err error) {
	cursor, err := c.Collection().Find(ctx, filter, opts...)
	if err != nil {
		return
	}
	err = cursor.Decode(result)
	if err != nil {
		return
	}
	return
}

func FindOne(c CollectionInterface, result interface{}, ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) (err error) {
	err = c.Collection().FindOne(ctx, filter, opts...).Decode(result)
	if err != nil {
		return
	}
	return
}
