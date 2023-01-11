package database

import "go.mongodb.org/mongo-driver/mongo"

// dummy cluster on mongodb atlas
var connectionString = "mongodb+srv://folafunmi:XpZaCsUw08AZgCWz@cluster0.moltryd.mongodb.net/?retryWrites=true&w=majority"

type DB struct {
	client *mongo.Client
}

func Connect() *DB {
	mongo.NewClient()
}
