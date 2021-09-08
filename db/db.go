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

func SearchProfile(query string) (profile.Profile, error) {
	client := connect()
	defer client.Disconnect(ctx)
	defer cancel()
	var res profile.Profile

	profilesClt := client.Database("hub").Collection("profiles")

	cursor, err := profilesClt.Find(ctx, bson.M{})
	check("Find: ", err)

	var profiles []bson.M

	err = cursor.All(ctx, &profiles)
	check("All results: ", err)

	for _, perfil := range profiles {
		if perfil["fullname"] != query {
			continue
		}

		res = profile.Profile{
			FullName: fmt.Sprint(perfil["fullname"]),
			Age: fmt.Sprint(perfil["age"]),
			Corporation: fmt.Sprint(perfil["corp"]),
			Experience: fmt.Sprint(perfil["exp"]),
			LinkedIn: fmt.Sprint(perfil["lkin"]),
			Twitter: fmt.Sprint(perfil["tw"]),
			Facebook: fmt.Sprint(perfil["fb"]),
			Instagram: fmt.Sprint(perfil["ig"]),
			Autorization: fmt.Sprint(perfil["aut"]),
		}
	}

	return res, nil
}

func AddProfile(p profile.Profile) error {
	client := connect()
	defer client.Disconnect(ctx)
	defer cancel()

	profilesClt := client.Database("hub").Collection("profiles")

	if p.Autorization != "yes" {
		return fmt.Errorf("Client not autorizes the public access for the his informations")
	}

	_, err := profilesClt.InsertOne(ctx, bson.D{
		{Key: "fullname", Value: p.FullName},
		{Key: "age", Value: p.Age},
		{Key: "corp", Value: p.Corporation},
		{Key: "exp", Value: p.Experience},
		{Key: "langs", Value: p.Languages},
		{Key: "lkin", Value: p.LinkedIn},
		{Key: "tw", Value: p.Twitter},
		{Key: "fb", Value: p.Facebook},
		{Key: "ig", Value: p.Instagram},
		{Key: "aut", Value: p.Autorization},
	})

	check("Insert into: ", err)
	return nil
}

func AllProfiles() []profile.Profile {
	client := connect()
	defer client.Disconnect(ctx)
	defer cancel()

	var bsonProfiles []bson.M
	var profiles []profile.Profile

	profilesClt := client.Database("hub").Collection("profiles")
	cursor, err := profilesClt.Find(ctx, bson.M{})
	check("Find all profiles: ", err)

	err = cursor.All(ctx, &bsonProfiles)
	check("All find: ", err)

	for _, perfil := range bsonProfiles {
		profiles = append(profiles, profile.Profile{
			FullName: fmt.Sprint(perfil["fullname"]),
			Age: fmt.Sprint(perfil["age"]),
			Corporation: fmt.Sprint(perfil["corp"]),
			Experience: fmt.Sprint(perfil["exp"]),
			Languages: fmt.Sprint(perfil["langs"]),
			LinkedIn: fmt.Sprint(perfil["lkin"]),
			Twitter: fmt.Sprint(perfil["tw"]),
			Facebook: fmt.Sprint(perfil["fb"]),
			Instagram: fmt.Sprint(perfil["ig"]),
			Autorization: fmt.Sprint(perfil["aut"]),
		})
	}

	return profiles
}
