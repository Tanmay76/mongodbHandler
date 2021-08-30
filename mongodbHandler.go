package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Disconnet(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	defer cancel()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func Connect(uri string) (*mongo.Client, context.Context, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
}

func Ping(client *mongo.Client, ctx context.Context) error {
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	fmt.Println("connected successfully")
	return nil
}

func Insertdata(blogUser *Bloguser) (primitive.ObjectID, error) {
	client, ctx, cancel, err := Connect(URLD)
	defer Disconnet(client, ctx, cancel)
	Ping(client, ctx)
	if err != nil {

		log.Printf("There is a problem in a connection %v", err)
	}

	var idr *mongo.InsertOneResult

	idr, err = client.Database(DATABASE).Collection(COLLECTION).InsertOne(ctx, blogUser)
	return idr.InsertedID.(primitive.ObjectID), err
}

func GetAlldate() ([]*Bloguser, error) {
	client, ctx, cancel, err := Connect(URLD)
	defer Disconnet(client, ctx, cancel)
	Ping(client, ctx)
	if err != nil {

		log.Printf("There is a problem in a connection %v", err)
	}

	filter := bson.M{}
	var blogu []*Bloguser
	cur, err := client.Database(DATABASE).Collection(COLLECTION).Find(ctx, filter)
	if err != nil {
		log.Printf("Cursor error:%v", err)
		return nil, err
	}
	defer cur.Close(ctx)
	err = cur.All(ctx, &blogu)
	return blogu, err

}

func FindOneData(id int64) (*Bloguser, error) {
	var blgou Bloguser
	filter := bson.M{"Id": id}
	client, ctx, cancel, err := Connect(URLD)
	if err != nil {
		log.Printf("There is a problem in a connection %v", err)
	}
	defer Disconnet(client, ctx, cancel)
	_ = Ping(client, ctx)
	err = client.Database(DATABASE).Collection(COLLECTION).FindOne(ctx, filter).Decode(&blgou)
	return &blgou, err
}

func FindOnebyemail(email string) (*Bloguser, error) {
	var blgou Bloguser
	filter := bson.M{"Email": email}
	client, ctx, cancel, err := Connect(URLD)
	if err != nil {
		log.Printf("There is a problem in a connection %v", err)
	}
	defer Disconnet(client, ctx, cancel)
	_ = Ping(client, ctx)
	err = client.Database(DATABASE).Collection(COLLECTION).FindOne(ctx, filter).Decode(&blgou)
	return &blgou, err
}
