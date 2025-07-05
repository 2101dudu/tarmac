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

type DBClient struct {
	Addr     string
	Username string
	Password string
}

type Service struct {
	dbClient    *DBClient
	mongoClient *mongo.Client
}

func (dbClient *DBClient) NewDbService() *Service {
	clientOptions := options.Client().ApplyURI("mongodb://" + dbClient.Addr)

	clientOptions.SetAuth(options.Credential{
		Username:   dbClient.Username,
		Password:   dbClient.Password,
		AuthSource: "admin",
	})

	mongoclient, err := mongo.Connect(clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	err = mongoclient.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Mongo ping failed:", err)
		return nil
	}

	return &Service{
		dbClient:    dbClient,
		mongoClient: mongoclient,
	}
}

func (db *Service) returnCollectionPointer(collection string) *mongo.Collection {
	return db.mongoClient.Database("tarmac").Collection(collection)
}

func CheckDBHit[T any](dbService *Service, collectionName string, id string) (*T, bool) {
	// defer logger.Log.TrackTime()()

	collection := dbService.returnCollectionPointer(collectionName)
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

func (db *Service) RefreshDB(collectionName string, id string, data any) {
	// defer logger.Log.TrackTime()()
	collection := db.returnCollectionPointer(collectionName)
	err := storeData(collection, id, data)
	if err != nil {
		logger.Log.Log("Failed DB refresh: ", err)
	} else {
		logger.Log.Log("DB refreshed for: ", collectionName)
	}
}

func GetAllProducts[T any](dbService *Service, collectionName string) ([]*T, error) {
	// defer logger.Log.TrackTime()()
	collection := dbService.returnCollectionPointer(collectionName)
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

func (db *Service) DeleteByID(collectionName string, id string) error {
	collection := db.returnCollectionPointer(collectionName)
	_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	return err
}

func (db *Service) RemoveCollections() {
	collections, err := db.mongoClient.Database("tarmac").ListCollectionNames(context.TODO(), struct{}{})
	if err != nil {
		logger.Log.Log("Failed to list Mongo collections:", err)
		return
	}
	for _, col := range collections {
		if err := db.mongoClient.Database("tarmac").Collection(col).Drop(context.TODO()); err != nil {
			logger.Log.Log("Failed to drop collection "+col+":", err)
		} else {
			logger.Log.Log("Dropped collection:", col)
		}
	}
}
