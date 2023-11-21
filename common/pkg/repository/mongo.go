package repository

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInterface interface {
	Close() error
	Create(data interface{}) error
	Read(id interface{}, result interface{}) error
	Update(id interface{}, data interface{}) error
	Delete(id interface{}) error
}

// MongoDB is a generic MongoDB repository.
type MongoDB struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

// NewMongoDB creates a new MongoDB instance.
func NewMongoDB(connectionString, databaseName, collectionName string, ctx context.Context) (MongoInterface, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return &MongoDB{}, err
	}

	db := client.Database(databaseName)
	col := db.Collection(collectionName)

	return &MongoDB{
		client:     client,
		database:   db,
		collection: col,
	}, nil
}

// Close closes the MongoDB client connection.
func (db *MongoDB) Close() error {
	if db.client != nil {
		return db.client.Disconnect(context.Background())
	}
	return nil
}

// Create inserts a document into the collection.
func (db *MongoDB) Create(data interface{}) error {
	_, err := db.collection.InsertOne(context.Background(), data)
	return err
}

// Read retrieves a document by ID from the collection.
func (db *MongoDB) Read(id interface{}, result interface{}) error {
	filter := primitive.E{Key: "_id", Value: id}
	err := db.collection.FindOne(context.Background(), filter).Decode(result)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil // Document not found
	}
	return err
}

// Update updates a document in the collection.
func (db *MongoDB) Update(id interface{}, data interface{}) error {
	filter := primitive.E{Key: "_id", Value: id}
	update := primitive.E{Key: "$set", Value: data}
	_, err := db.collection.UpdateOne(context.Background(), filter, update)
	return err
}

// Delete removes a document by ID from the collection.
func (db *MongoDB) Delete(id interface{}) error {
	filter := primitive.E{Key: "_id", Value: id}
	_, err := db.collection.DeleteOne(context.Background(), filter)
	return err
}
