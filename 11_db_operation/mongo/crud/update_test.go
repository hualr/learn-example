package crud

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"testing"
)

//更新多条文档
func UpdateManyUsers(client *mongo.Client) {
	coll := client.Database("tiktok").Collection("user")
	filter := bson.D{{"user_id", bson.D{{"$gt", 0}}}}
	update := bson.D{{"$mul", bson.D{{"FollowCount", 1}}}}
	result, err := coll.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

//查询到匹配数据则修改，否则就插入
func UpdateOrInsert(client *mongo.Client) {
	coll := client.Database("tiktok").Collection("user")
	filter := bson.D{{"username", "Amy"}}
	update := bson.D{{"$set", bson.D{{"username", "Susan"}}}}
	opts := options.Update().SetUpsert(true)
	result, err := coll.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Number of documents updated: %v\n", result.ModifiedCount)
	fmt.Printf("Number of documents upserted: %v\n", result.UpsertedCount)
}

func Test3(t *testing.T) {
	client, err := ConnectToDB("mongodb://101.201.67.114:27017", "test", 3, 64)

	if err != nil {
		log.Printf("mongo connect error")
		return
	}
	if client != nil {
		UpdateManyUsers(client)
	}
}
