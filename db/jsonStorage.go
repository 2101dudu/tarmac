package db

import (
	"context"
	"tarmac/json"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type SOAPStore struct {
	Payload   []byte    `bson:"payload"`
	FetchedAt time.Time `bson:"fetched_at"`
}

func storeData(collection *mongo.Collection, id string, data any) error {
	// defer logger.Log.TrackTime()()
	compressedData, err := json.Compress(data)
	if err != nil {
		return err
	}

	doc := bson.M{
		"_id":        id,
		"payload":    compressedData,
		"fetched_at": time.Now(),
	}

	_, err = collection.ReplaceOne(
		context.Background(),
		bson.M{"_id": id},
		doc,
		options.Replace().SetUpsert(true),
	)

	return err
}

func loadData(collection *mongo.Collection, id string) (SOAPStore, error) {
	// defer logger.Log.TrackTime()()
	var result SOAPStore
	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&result)
	return result, err
}
