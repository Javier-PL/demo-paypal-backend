package main

import (
	"os"
	"paypal-starter-backend2/m/database"
	"paypal-starter-backend2/m/routers"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"context"
	"time"

	"net/http"

	"github.com/Ramso-dev/env"

	"github.com/Ramso-dev/log"

	"github.com/rs/cors"
)

var Log log.Logger

func main() {

	var err error

	type Configuration struct {
		DB_CONNECT string
		CLIENT_ID  string
		SECRET     string
	}
	var config Configuration
	env.InitEnvVars(config)

	database.DBCon, err = mongo.NewClient(options.Client().ApplyURI(os.Getenv("DB_CONNECT")))
	if err != nil {
		Log.Error("Client:", err)
	}
	client := database.DBCon

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		Log.Error("Connect:", err)
	} else {

		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		err = client.Ping(ctx, readpref.Primary())

		if err != nil {
			Log.Error("Connect:", err)
		}
		Log.Info("Connected successfully")

	}

	router := routers.InitRoutes()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8082" // Default port if not specified
	}

	Log.Info("Starting http.ListenAndServe service on port", port)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"*"},
	})

	handler := c.Handler(router)

	http.ListenAndServe(":"+port, handler)

}
