package dbhelper

import (
	"context"
	"fmt"
	"log"

	"github.com/suvam720/mongoapi/pkg/database"
	"github.com/suvam720/mongoapi/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertTask(tosk model.Task) {
	inserted, err := database.Callection.InsertOne(context.Background(), tosk)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(inserted.InsertedID)
}

func GetTasks() []primitive.M {
	cur, err := database.Callection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	var tasks []primitive.M
	for cur.Next(context.Background()) {
		var task bson.M
		err := cur.Decode(&task)
		if err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, task)
	}
	defer cur.Close(context.Background())

	return tasks
}

func UpdateTask(toskId string) {
	id, err := primitive.ObjectIDFromHex(toskId)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"completed": true}}

	res, err := database.Callection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res.ModifiedCount)

}

func DeleteTask(toskId string) {
	id, _ := primitive.ObjectIDFromHex(toskId)
	filter := bson.M{"_id": id}

	delCount, err := database.Callection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(delCount.DeletedCount)
}

func DeleteTasks() int64 {
	deleteCount, err := database.Callection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	
	return deleteCount.DeletedCount
}
