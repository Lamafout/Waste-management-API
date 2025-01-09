package model

import (
	"context"
	"log"
	"time"
	"waste_management/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	log.Println("Connection is created!")

	return &connection{Client: *client, config: config}
}

func (c *connection) Create(collection string, document interface{}) error {
	_, err := c.Client.Database(c.config.Database).Collection(collection).InsertOne(context.Background(), document)
	return err
}

func (c *connection) Read(collection string, id string) (map[string]interface{}, error) {
	bsonId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Println("Error while converting id into bson type id: ", err.Error())
	}

	filter := bson.M{"_id": bsonId}
	result := map[string]interface{}{}
	err = c.Client.Database(c.config.Database).Collection(collection).FindOne(context.Background(), filter).Decode(&result)

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

func (c *connection) ReadFiltered(collection string, clientFilter string, filedName string) ([]map[string]interface{}, error) {
	filter := bson.M{filedName:clientFilter}

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