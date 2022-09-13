package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
	"time"
)

type TimePorint struct {
	StartTime int64 `bson:"startTime"` //开始时间
	EndTime   int64 `bson:"endTime"`   //结束时间
}
type LogRecord struct {
	JobName string     `bson:"jobName"` //任务名
	Command string     `bson:"command"` //shell命令
	Err     string     `bson:"err"`     //脚本错误
	Content string     `bson:"content"` //脚本输出
	Tp      TimePorint //执行时间
}

func ConnectToDB(uri, name string, timeout time.Duration, num uint64) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	o := options.Client().ApplyURI(uri)
	o.SetMaxPoolSize(num)
	client, err := mongo.Connect(ctx, o)
	if err != nil {
		return nil, err
	}

	return client.Database(name), nil
}

func Test2(t *testing.T) {
	var (
		client     *mongo.Client
		err        error
		db         *mongo.Database
		collection *mongo.Collection
	)
	//1.建立连接
	if client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://101.201.67.114:27017").SetConnectTimeout(5*time.Second)); err != nil {
		fmt.Print(err)
		return
	}
	//2.选择数据库 my_db
	db = client.Database("my_db")

	//3.选择表 my_collection
	collection = db.Collection("my_collection")
	collection = collection

	//批量插入
	collection.InsertMany(context.TODO(), []interface{}{
		LogRecord{
			JobName: "job10",
			Command: "echo 1",
			Err:     "",
			Content: "1",
			Tp: TimePorint{
				StartTime: time.Now().Unix(),
				EndTime:   time.Now().Unix() + 10,
			},
		},
		LogRecord{
			JobName: "job10",
			Command: "echo 2",
			Err:     "",
			Content: "2",
			Tp: TimePorint{
				StartTime: time.Now().Unix(),
				EndTime:   time.Now().Unix() + 10,
			},
		},
		LogRecord{
			JobName: "job10",
			Command: "echo 3",
			Err:     "",
			Content: "3",
			Tp: TimePorint{
				StartTime: time.Now().Unix(),
				EndTime:   time.Now().Unix() + 10,
			},
		},
		LogRecord{
			JobName: "job10",
			Command: "echo 4",
			Err:     "",
			Content: "4",
			Tp: TimePorint{
				StartTime: time.Now().Unix(),
				EndTime:   time.Now().Unix() + 10,
			},
		},
	})

}
