package main

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func findDocument(client *mongo.Client, ctx context.Context, database, collectionName string, query interface{}) *mongo.Cursor {

	collection := client.Database(database).Collection(collectionName)

	result, err := collection.Find(ctx, query)

	check(err)

	return result
}

func insertDocument(client *mongo.Client, ctx context.Context, database, collectionName string, document interface{}) *mongo.InsertOneResult {

	collection := client.Database(database).Collection(collectionName)

	result, err := collection.InsertOne(ctx, document)

	check(err)

	return result
}

func updateDocument(client *mongo.Client, ctx context.Context, database, collectionName string, _id, document interface{}) *mongo.UpdateResult {
	collection := client.Database(database).Collection(collectionName)

	result, err := collection.UpdateOne(ctx, _id, document)

	check(err)

	return result
}

func deleteDocument(client *mongo.Client, ctx context.Context, database, collectionName string, _id interface{}) *mongo.DeleteResult {
	collection := client.Database(database).Collection(collectionName)

	result, err := collection.DeleteOne(ctx, _id)

	check(err)

	return result
}
