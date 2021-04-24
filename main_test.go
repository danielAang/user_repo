package main_test

import (
	"context"
	"log"
	"os"
	"testing"

	main "github.com/danielAang/user_repo"
)

var a main.App

func TestMain(m *testing.M) {
	a.Initialize(
		os.Getenv("APP_DB_NAME"),
		os.Getenv("APP_DB_URI"),
	)
	ensurePingSuccessful()
}

func ensurePingSuccessful() {
	if err := a.DB.Client().Ping(context.TODO(), nil); err != nil {
		log.Fatal(err)
	}
}
