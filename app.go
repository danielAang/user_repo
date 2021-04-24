package main

import (
	"context"
	"log"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type App struct {
	Router *mux.Router
	DB     *mongo.Database
}

func (a *App) Initialize(dbName, uri string) {
	ctx := context.TODO()
	co := options.Client().ApplyURI(uri)
	c, err := mongo.Connect(ctx, co)
	if err != nil {
		log.Fatal(err)
	}
	c.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	a.DB = c.Database(dbName)
}

func (a *App) Run(addr string) {}
