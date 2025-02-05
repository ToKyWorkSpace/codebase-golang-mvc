package mongodb

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"codebase-golang/internal/app/config"
	Models "codebase-golang/internal/app/model"
	utils "codebase-golang/internal/app/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	mongoClient    *mongo.Client
	postCollection *mongo.Collection
	ctx            context.Context
)

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbUri := config.GetEnv("MONGO_URL")
	connectionOpts := options.Client().ApplyURI(dbUri)
	mongoClient, err := mongo.Connect(ctx, connectionOpts)
	if err != nil {
		fmt.Printf("an error ocurred when connect to mongoDB : %v", err)
		panic(err)
	}

	if err = mongoClient.Ping(ctx, readpref.Primary()); err != nil {
		fmt.Printf("an error ocurred when connect to mongoDB : %v", err)
		panic(err)
	}

	postCollection = mongoClient.
		Database(config.GetEnv("MONGO_DATABASE")).
		Collection(config.GetEnv("MONGO_COLLECTION"))
}

func GetAll(c *gin.Context) []*Models.Users {
	query := bson.M{}
	cursor, err := postCollection.Find(ctx, query)
	if err != nil {
		fmt.Printf("an error ocurred when get data from mongoDB : %v", err)
		return nil
	}
	defer cursor.Close(ctx)

	var users []*Models.Users
	for cursor.Next(ctx) {
		user := Models.Users{}
		err = cursor.Decode(&user)

		if err != nil {
			fmt.Printf("an error ocurred when parse data from cursor : %v", err)
			return nil
		}

		users = append(users, &user)
	}
	return users
}

func Get(c *gin.Context) *Models.Users {
	param := c.Param("id")
	id, _ := primitive.ObjectIDFromHex(param)
	query := bson.M{"_id": id}

	cursor, err := postCollection.Find(ctx, query)
	if err != nil {
		fmt.Printf("an error ocurred when get data from mongoDB : %v", err)
		return nil
	}
	defer cursor.Close(ctx)

	var users []*Models.Users
	for cursor.Next(ctx) {
		user := Models.Users{}
		err = cursor.Decode(&user)

		if err != nil {
			fmt.Printf("an error ocurred when parse data from cursor : %v", err)
			return nil
		}
		users = append(users, &user)
	}
	return users[0]
}

func Add(c *gin.Context) *Models.PostUsers {
	age, _ := strconv.Atoi(c.PostForm("age"))
	payload := Models.PostUsers{
		Username: c.PostForm("username"),
		Password: utils.Encrypt(c.PostForm("password")),
		Name:     c.PostForm("name"),
		Age:      age,
	}

	_, err := postCollection.InsertOne(ctx, payload)

	if err != nil {
		fmt.Printf("an error ocurred when connect to mongoDB : %v", err)
		return nil
	}

	return &payload
}

func Update(c *gin.Context) *Models.PostUsers {
	param := c.Param("id")
	id, _ := primitive.ObjectIDFromHex(param)
	age, _ := strconv.Atoi(c.PostForm("age"))
	query := bson.M{"_id": id}

	payload := Models.PostUsers{
		Username: c.PostForm("username"),
		Password: utils.Encrypt(c.PostForm("password")),
		Name:     c.PostForm("name"),
		Age:      age,
	}
	command := bson.D{{Key: "$set", Value: payload}}

	res := postCollection.FindOneAndUpdate(ctx, query, command)
	var result *Models.Users
	if err := res.Decode(&result); err != nil {
		fmt.Printf("an error ocurred when update data : %v", err)
		return nil
	}

	return &payload
}

func Delete(c *gin.Context) int64 {
	param := c.Param("id")
	objectID, _ := primitive.ObjectIDFromHex(param)
	query := bson.D{{Key: "_id", Value: objectID}}
	res, err := postCollection.DeleteOne(ctx, query)

	if err != nil {
		fmt.Printf("an error ocurred when delete data : %v", err)
		panic(err)
	}
	println(res.DeletedCount)
	return res.DeletedCount

}
