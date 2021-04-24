package main_test

import (
	"bytes"
	"context"
	"encoding/json"
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
	resp := executeRequest(req)
	checkResponseCode(t, http.StatusOK, resp.Code)
	if body := resp.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s instead.", body)
	}
}

func TestNonExistentUser(t *testing.T) {
	req, _ := http.NewRequest("GET", "/users/abc", nil)
	resp := executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, resp.Code)
	var m map[string]string
	json.Unmarshal(resp.Body.Bytes(), &m)
	if m["error"] != "User not found" {
		t.Errorf("Expected key 'error' to be 'User not found'")
	}
}

func TestCreateUser(t *testing.T) {
	jsonUser := []byte(`{ "name": "john doe", "email": "johndoe@email.com", "password":"123", "phones": ["12345678"] }`)
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonUser))
	req.Header.Set("Content-Type", "application/json")
	resp := executeRequest(req)
	checkResponseCode(t, http.StatusCreated, resp.Code)
	var m map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &m)
	if m["id"] != nil {
		t.Errorf("Expected id to don't be nil.")
	}
	if m["name"] != "jonh doe" {
		t.Errorf("Expected name to be 'john doe'. Got %s instead.", m["name"])
	}
	if m["email"] != "jonh doe" {
		t.Errorf("Expected email to be 'john doe'. Got %s instead.", m["email"])
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rc := httptest.NewRecorder()
	a.Router.ServeHTTP(rc, req)
	return rc
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d", http.StatusOK, actual)
	}
}
