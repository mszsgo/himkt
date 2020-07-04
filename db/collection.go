package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Find(c *mongo.Collection, result interface{}, ctx context.Context, filter interface{}, opts ...*options.FindOptions) (err error) {
	cursor, err := c.Find(ctx, filter, opts...)
	if err != nil {
		return err
	}
	if cursor.Err() != nil {
		return cursor.Err()
	}
	err = cursor.All(ctx, result)
	if err != nil {
		return err
	}
	return
}

func Page(c *mongo.Collection, result interface{}, ctx context.Context, filter interface{}, opts ...*options.FindOptions) (total int64, err error) {
	err = Find(c, result, ctx, filter, opts...)
	if err != nil {
		return 0, err
	}
	return c.CountDocuments(ctx, filter)
}
