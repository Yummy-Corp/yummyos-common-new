package config

import (
	"context"
	"log"
	"net/url"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func InitMongoDB() *mongo.Client {
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	cluster := os.Getenv("DB_INSTANCE")
	authSource := os.Getenv("DB_AUTHSOURCE")
	authMechanism := os.Getenv("DB_AUTHMECHANISM")

	uri := "mongodb+srv://" + url.QueryEscape(username) + ":" +
		url.QueryEscape(password) + "@" + cluster +
		"/?authSource=" + authSource +
		"&authMechanism=" + authMechanism + "&retryWrites=true&w=majority"

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPIOptions)
	db, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err := db.Connect(ctx); err != nil {
		log.Fatal(err)
		return nil
	}

	if err := db.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalf("Error pinging db: %s", err)
		return nil
	}

	return db
}
