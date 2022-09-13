package crud

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"testing"
)

func DeleteOneVideo(Client *mongo.Client) {
	coll := Client.Database("tiktok").Collection("video")
	filter := bson.D{{"video_id", 10}}
	result, err := coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

//删除多条文档
func DeleteVideos(Client *mongo.Client) {
	coll := Client.Database("tiktok").Collection("video")
	filter := bson.D{{"VideoId", bson.D{{"$gt", 9}}}}
	results, err := coll.DeleteMany(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	fmt.Println(results)
}

func Test2(t *testing.T) {
	client, err := ConnectToDB("mongodb://101.201.67.114:27017", "test", 3, 64)

	if err != nil {
		log.Printf("mongo connect error")
		return
	}
	if client != nil {
		DeleteOneVideo(client)
	}
}
