package database

import (
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

// dummy cluster on mongodb atlas
var connectionString = os.Getenv("BASE_URL")

type DB struct {
	client *mongo.Client
}

func Connect() *DB {
	mongo.NewClient()
}
