package main

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Phones    []string  `json:"phones"`
	Birth     time.Time `json:"birthDate"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (u *User) Create(db *mongo.Database) error {
	return errors.New("not implemented")
}

func (u *User) Update(db *mongo.Database) error {
	return errors.New("not implemented")
}

func (u *User) Delete(db *mongo.Database) error {
	return errors.New("not implemented")
}

func (u *User) FindById(id string, db *mongo.Database) error {
	return errors.New("not implemented")
}

func (u *User) FindAll(from, to int, db *mongo.Database) ([]User, error) {
	return nil, errors.New("not implemented")
}
