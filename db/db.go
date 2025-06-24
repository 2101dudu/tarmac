package db

import (
	"context"
	"log"
	"tarmac/json"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Client struct {
	Addr     string
	Username string
	Password string
}

func Start(dbClient *Client) *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://" + dbClient.Addr)

	clientOptions.SetAuth(options.Credential{
		Username:   dbClient.Username,
		Password:   dbClient.Password,
		AuthSource: "admin",
	})

	client, err := mongo.Connect(clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Mongo ping failed:", err)
		return nil
	}

	return client
}

func returnCollectionPointer(dbClient *mongo.Client, collection string) *mongo.Collection {
	return dbClient.Database("tarmac").Collection(collection)
}

func CheckDBHit[T any](dbClient *mongo.Client, collectionName string, id string) (*T, bool) {
	collection := returnCollectionPointer(dbClient, collectionName)
	soapStore, err := loadData(collection, collectionName, id)
	if err == mongo.ErrNoDocuments {
		log.Println("No DB Hit for", collectionName)
		return nil, false
	}
	log.Println("DB Hit for", collectionName)
	uncompressedData, err := json.Uncompress[T](soapStore.Payload)
	if err != nil {
		log.Println("Error uncompressing json:", err)
		return nil, false
	}
	return uncompressedData, time.Since(soapStore.FetchedAt) > 24*time.Hour
}

func RefreshDB(dbClient *mongo.Client, collectionName string, id string, data any) {
	collection := returnCollectionPointer(dbClient, collectionName)
	err := storeData(collection, collectionName, id, data)
	if err != nil {
		log.Println("Failed DB refresh:", err)
	} else {
		log.Println("DB refreshed for", collectionName)
	}
}
