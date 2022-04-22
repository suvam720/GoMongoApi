package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://admin:admin@cluster0.dmqyx.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
const dbName = "todo"
const colName = "task"

var Callection *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection Sucessful")
	Callection = client.Database(dbName).Collection(colName)
}
