package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/ahub-tech/hub-profile-api/profile"
	"github.com/henriquetied472/logplus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const msc = "mongodb+srv://hubdbroot:ISQ0zawYigoejYYt@cluster0.umlw4.mongodb.net/Cluster0?retryWrites=true&w=majority"

var cancel context.CancelFunc
var ctx context.Context

func check(msg string, err error) {
	if err != nil {
		logplus.Fatal(msg + err.Error())
		os.Exit(1)
	}
}

func connect() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(msc))
	check("Client create: ", err)

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	check("Connect db: ", err)

	err = client.Ping(ctx, nil)
	check("Ping db: ", err)

	logplus.Debug("Connected to mongodb")

	return client
}

func SearchProfile(query string) {
	client := connect()
	defer client.Disconnect(ctx)
	defer cancel()

	profilesClt := client.Database("hub").Collection("profiles")

	cursor, err := profilesClt.Find(ctx, bson.M{})
	check("Find: ", err)

	var profiles []bson.M

	err = cursor.All(ctx, &profiles)
	check("All results: ", err)

	for _, profile := range profiles {
		logplus.Debug(fmt.Sprint(profile))
	}
}

func AddProfile(p profile.Profile) {
	client := connect()
	defer client.Disconnect(ctx)
	defer cancel()

	hubdb := client.Database("hub")
	profilesClt := hubdb.Collection("profiles")

	_, err := profilesClt.InsertOne(ctx, bson.D{
		{Key: "fullname", Value: p.FullName},
		{Key: "age", Value: p.Age},
		{Key: "corp", Value: p.Corporation},
		{Key: "exp", Value: p.Experience},
		{Key: "lkin", Value: p.LinkedIn},
		{Key: "tw", Value: p.Twitter},
		{Key: "fb", Value: p.Facebook},
		{Key: "ig", Value: p.Instagram},
		{Key: "aut", Value: p.Autorization},
	})

	check("Insert into: ", err)
}
