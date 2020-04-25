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
	if cursor.Next(ctx) {
		err = cursor.All(ctx, result)
		if err != nil {
			return err
		}
	}
	return
}
