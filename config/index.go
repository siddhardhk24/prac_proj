package config

import (
	"context"
	
	"log"
	"prac_proj/constants"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectDataBase() (*mongo.Client, error){
	ctx,_ := context.WithTimeout(context.Background(),10*time.Second)
	mongoConnection := options.Client().ApplyURI(constants.ConnectionString)
	mongoClient,err := mongo.Connect(ctx,mongoConnection)
	if err != nil{
		log.Fatal(err.Error())
		return nil, err
	}
	if err := mongoClient.Ping(ctx, readpref.Primary()); err!=nil{
		return nil, err
	}
	return mongoClient, nil
}

func GetCollection(client *mongo.Client,dbname string, collectionName string) *mongo.Collection{
	
	collection := client.Database(dbname).Collection(collectionName)
	return collection
}

