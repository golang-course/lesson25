package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {
	//connection
	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://root:secretpassword@localhost:27018/?authSource=admin"))
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Println(err)
			return
		}
	}()
	cursor := client.Database("users").Collection("users")

	//Document Count
	// count, err := cursor.CountDocuments(context.TODO(), bson.M{})
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// fmt.Println(count)

	// var Dmitry map[string]any
	// err = cursor.FindOne(context.TODO(), bson.M{"name": "Dmitry"}).Decode(&Dmitry)
	// if err != nil {
	// 	if err == mongo.ErrNoDocuments {
	// 		log.Println("No such document")
	// 	} else {
	// 		log.Println(err)
	// 		return
	// 	}
	// } else {
	// 	fmt.Println(Dmitry)
	// }

	// _, err = cursor.InsertOne(context.TODO(), bson.M{"name": "Serega", "work": "CPO"})
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// var Serega []map[string]any
	//  err = cursor.FindOne(context.TODO(), bson.M{"name": "Serega"}).Decode(&Serega)
	// if err != nil {
	// 	if err == mongo.ErrNoDocuments {
	// 		log.Println("No such document")
	// 	} else {
	// 		log.Println(err)
	// 		return
	// 	}
	// } else {
	// 	fmt.Println(Serega)
	// }

	cur, err := cursor.Find(context.TODO(), bson.M{"name": "Serega"})
	if err != nil {
		log.Println(err)
		return
	}
	defer cur.Close(context.TODO())

	var res []map[string]any
	err = cur.All(context.TODO(), &res)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
