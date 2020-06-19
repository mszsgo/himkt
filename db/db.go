package db

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
	"time"
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

var _collection_tx = make(map[string]bool)

func Collection(uri, name string) *mongo.Collection {
	c := Connect(uri).Collection(name)

	// 第一次加载判断集合中是否存在数据，因事务操作不允许空集合
	if !_collection_tx[name] {
		count, _ := c.CountDocuments(nil, bson.M{}, options.Count().SetLimit(1))
		if count == 0 {
			c.InsertOne(nil, bson.M{"createdAt": time.Now()})
		}
		_collection_tx[name] = true
	}
	return c
}
