package db

import (
	"context"
	"github.com/a9u/we-mock-api/config"
	"github.com/a9u/we-mock-api/pkg/wlog"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

func InitDb(conf *config.Conf) *mongo.Database {
	uri := conf.Database.Uri()

	client, err := newDbConnection(uri)
	if err != nil {
		wlog.Error("Failed to connect to db", nil)
		panic(err)
	}

	wlog.Info("successfully connected to db")
	return client.Database(conf.Database.Name)
}

func newDbConnection(uri string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		wlog.Error("connection to mongodb failed", err)
		return nil, err
	}

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		wlog.Error("ping failed", err)
		return nil, err
	}

	return client, nil
}
