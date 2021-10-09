package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Post struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserId  string             `json:"userId,omitempty" bson:"userId,omitempty"`
	Caption string             `json:"caption,omitempty" bson:"caption,omitempty"`
	Iurl    string             `json:"iurl,omitempty" bson:"iurl,omitempty"`
	Tstamp  string             `json:"tstamp,omitempty" bson:"tstamp,omitempty"`
}

const collectionName = "post"

func CreatePostEndpoint(response http.ResponseWriter, request *http.Request) {
	Time := time.Date(2020, 11, 14, 10, 45, 16, 0, time.UTC)
	t := Time.String()
	response.Header().Set("content-type", "application/json")
	var post Post
	_ = json.NewDecoder(request.Body).Decode(&post)
	post.Tstamp = t
	collection := client.Database("apointi").Collection(collectionName)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, post)
	json.NewEncoder(response).Encode(result)
}

func GetPostById(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var post Post
	collection := client.Database("apointi").Collection(collectionName)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := collection.FindOne(ctx, Post{ID: id}).Decode(&post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(post)
}

func GetPostByUserEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := params["id"]
	var post []Post
	page, _ := strconv.Atoi(params["pages"])
	var perPage int64 = 2

		


	collection := client.Database("TASK1").Collection("post")
	

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	filter := bson.M{}
	collection.CountDocuments(ctx, filter)
	findOptions := options.Find()
	findOptions.SetSkip((int64(page) - 1) * perPage)
	findOptions.SetLimit(perPage)



	cursor, err := collection.Find(ctx, Post{UserId: id},findOptions)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var person Post
		cursor.Decode(&person)
		post = append(post, person)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(post)
}