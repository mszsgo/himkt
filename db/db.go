package db

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
)

var _db = make(map[string]*mongo.Database)

// 注意处理连接panic
func Connect(connectionString string) *mongo.Database {
	if _db[connectionString] != nil {
		return _db[connectionString]
	}

	if connectionString == "" {
		log.Panic(errors.New("Mongodb 连接字符串不能为空"))
	}
	dbName := (strings.Split((strings.Split(connectionString, "/"))[3], "?"))[0]
	if dbName == "" {
		log.Panic(errors.New(fmt.Sprintf("Errror Mongodb connectionString %s", connectionString)))
	}
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	client, err := mongo.Connect(nil, options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Panic(errors.New(fmt.Sprintf("Errror Connect mongodb exception %s", err)))
	}
	database := client.Database(dbName)
	names, err := client.ListDatabaseNames(nil, bson.D{})
	if err != nil {
		log.Panic(errors.New(fmt.Sprintf("ListDatabaseNames Connect exception: %s", err)))
	}
	log.Printf("Mongodb connect success -> %s  DatabaseNames->%s", connectionString, names)

	_db[connectionString] = database
	return database
}
