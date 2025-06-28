package db

import (
	"context"
	"log"
	"tarmac/json"
	"tarmac/logger"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
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
	defer logger.Log.TrackTime()()

	collection := returnCollectionPointer(dbClient, collectionName)
	soapStore, err := loadData(collection, id)
	if err == mongo.ErrNoDocuments {
		return nil, false
	}
	uncompressedData, err := json.Uncompress[T](soapStore.Payload)
	if err != nil {
		logger.Log.Log("Error uncompressing json:", err)
		return nil, false
	}
	return uncompressedData, time.Since(soapStore.FetchedAt) > 24*time.Hour
}

func RefreshDB(dbClient *mongo.Client, collectionName string, id string, data any) {
	defer logger.Log.TrackTime()()
	collection := returnCollectionPointer(dbClient, collectionName)
	err := storeData(collection, id, data)
	if err != nil {
		logger.Log.Log("Failed DB refresh: ", err)
	} else {
		logger.Log.Log("DB refreshed for: ", collectionName)
	}
}

func GetAllProducts[T any](dbClient *mongo.Client, collectionName string) ([]*T, error) {
	defer logger.Log.TrackTime()()
	collection := returnCollectionPointer(dbClient, collectionName)
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	type doc struct {
		ID        string    `bson:"_id"`
		Payload   []byte    `bson:"payload"`
		FetchedAt time.Time `bson:"fetched_at"`
	}
	var documents []*doc

	if err := cursor.All(context.TODO(), &documents); err != nil {
		return nil, err
	}

	var products []*T
	for _, document := range documents {
		uncompressed, err := json.Uncompress[T](document.Payload)
		if err != nil {
			logger.Log.Log("Failed to uncompress payload:", err)
			continue
		}
		products = append(products, uncompressed)
	}

	return products, nil
}
