// Package database
package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"Server/internal/config"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

type MongoDBClient struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func ConnectDB() (*MongoDBClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Client().
		ApplyURI(config.Env.DB.URI).
		SetAuth(options.Credential{
			Username: config.Env.DB.Username,
			Password: config.Env.DB.Password,
		})

	client, err := mongo.Connect(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	log.Println("Connected to MongoDB successfully")
	return &MongoDBClient{
		Client:   client,
		Database: client.Database(config.Env.DB.Name),
	}, nil
}

func (c *MongoDBClient) CreateConstraints() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true).SetName("email_unique"),
	}
	name, err := c.Database.Collection("users").Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		return fmt.Errorf("failed to create unique index: %v", name)
	}

	log.Printf("Successfully created unique index: %s\n", name)
	return nil
}

func (c *MongoDBClient) Disconnect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := c.Database.Client().Disconnect(ctx); err != nil {
		return fmt.Errorf("error disconnecting database: %v", err)
	}

	log.Println("Database disconnected successfully")
	return nil
}
