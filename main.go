package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

const saltSize = 20

func main() {
	fmt.Println("Starting the application...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	router := mux.NewRouter()

	router.HandleFunc("/users", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/users/{id}", GetPersonEndpoint).Methods("GET")

	router.HandleFunc("/posts", CreatePostEndpoint).Methods("POST")
	router.HandleFunc("/posts/{id}", GetPostById).Methods("GET")
	router.HandleFunc("/posts/users/{id}/{pages}", GetPostByUserEndpoint).Methods("GET")
	http.ListenAndServe(":12345", router)

}