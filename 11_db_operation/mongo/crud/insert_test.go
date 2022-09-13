package crud

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"testing"
	"time"
)

func insertOneVideo(Client *mongo.Client) {
	coll := Client.Database("tiktok").Collection("video")
	video := Video{
		VideoId:       10,
		UserId:        1,
		PlayUrl:       "https://www.w3schools.com/html/movie.mp4",
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		Favorites:     []int64{},
		PublishDate:   time.Now(),
		Title:         "video10",
	}
	result, err := coll.InsertOne(context.TODO(), video)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func Test1(t *testing.T) {
	client, err := ConnectToDB("mongodb://101.201.67.114:27017", "test", 3, 64)

	if err != nil {
		log.Printf("mongo connect error")
		return
	}
	if client != nil {
		insertOneVideo(client)
	}

}
