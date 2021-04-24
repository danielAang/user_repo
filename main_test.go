package main_test

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	main "github.com/danielAang/user_repo"
)

var a main.App

func TestInitialization(m *testing.T) {
	a.Initialize(
		os.Getenv("APP_DB_NAME"),
		os.Getenv("APP_DB_URI"),
	)
	if err := a.DB.Client().Ping(context.TODO(), nil); err != nil {
		log.Fatal(err)
	}
}

func TestEmptyCollection(t *testing.T) {
	fmt.Println("Testing empty collection")
	req, _ := http.NewRequest("GET", "/users", nil)
	recorder := httptest.NewRecorder()
	a.Router.ServeHTTP(recorder, req)
	if http.StatusOK == recorder.Code {
		t.Errorf("Expected response code %d. Got %d", http.StatusOK, recorder.Code)
	}
	if body := recorder.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s instead.", body)
	}
}
