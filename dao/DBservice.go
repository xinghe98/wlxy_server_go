package dao

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 通过构造函数连接数据库,需要做数据库操作的service，将此实例注入到结构体

func NewDBService(uri, dbName string) (*DBService, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("could not connect to MongoDB: %v", err)
	}

	db := client.Database(dbName)
	return &DBService{db: db}, nil
}

type DBService struct {
	db *mongo.Database
}

func (db *DBService) GetCollection(collectionName string) *mongo.Collection {
	return db.db.Collection(collectionName)
}
