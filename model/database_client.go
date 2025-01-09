package model

import (
	"context"
	"time"
	"waste_management/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type connection struct {
	Client mongo.Client
	config config.MongoDbConfig
}

func NewMongoDBConnection(config config.MongoDbConfig) *connection {
	clientOptions := options.Client().ApplyURI(config.Uri)
	context, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	client, err := mongo.Connect(context, clientOptions)

	if err != nil {
		panic(err)
	}

	if err = client.Ping(context, nil); err != nil {
		panic(err)
	}

	return &connection{Client: *client, config: config}
}

func (c *connection) Create(collection string, document interface{}) error {
	_, err := c.Client.Database(c.config.Database).Collection(collection).InsertOne(context.Background(), document)
	return err
}

func (c *connection) Read(collection string, id int64) (map[string]interface{}, error) {
	filter := bson.M{"_id": id}
	result := map[string]interface{}{}
	err := c.Client.Database(c.config.Database).Collection(collection).FindOne(context.Background(), filter).Decode(&result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *connection) ReadAll(collection string) ([]map[string]interface{}, error) {
    filter := bson.M{}

    cursor, err := c.Client.Database(c.config.Database).Collection(collection).Find(context.Background(), filter)
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.Background())

    var results []map[string]interface{}

    for cursor.Next(context.Background()) {
        var result map[string]interface{}
        if err := cursor.Decode(&result); err != nil {
            return nil, err
        }
        results = append(results, result)
    }

    if err := cursor.Err(); err != nil {
        return nil, err
    }

    return results, nil
}