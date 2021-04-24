package main

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const CollectionName = "users"

type User struct {
	ID        string    `json:"id" bson:"_id,omitempty"`
	Name      string    `json:"name" bson:"name,omitempty"`
	Email     string    `json:"email" bson:"email,omitempty"`
	Password  string    `json:"password" bson:"password,omitempty"`
	Phones    []string  `json:"phones" bson:"phones,omitempty"`
	Birth     time.Time `json:"birthDate" bson:"birthDate,omitempty"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt,omitempty"`
}

func (u *User) Create(db *mongo.Database) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	result, err := db.Collection(CollectionName).InsertOne(ctx, u)
	if err != nil {
		return nil, err
	}
	u.ID = result.InsertedID.(primitive.ObjectID).Hex()
	return u, nil
}

func (u *User) Update(db *mongo.Database) error {
	return errors.New("not implemented")
}

func (u *User) Delete(db *mongo.Database) error {
	return errors.New("not implemented")
}

func (u *User) FindById(db *mongo.Database) (*User, error) {
	var user User
	hex, err := primitive.ObjectIDFromHex(u.ID)
	if err != nil {
		return nil, errors.New("unable to parse id")
	}
	query := bson.M{"_id": hex}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor := db.Collection(CollectionName).FindOne(ctx, query)
	err = cursor.Decode(&user)
	if err != nil {
		return nil, errors.New("unable to parse users")
	}
	return &user, nil
}

func (u *User) FindAll(from, to int, db *mongo.Database) ([]User, error) {
	return nil, errors.New("not implemented")
}
