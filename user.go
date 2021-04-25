package main

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (u *User) Update(db *mongo.Database) (*User, error) {
	hex, err := primitive.ObjectIDFromHex(u.ID)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	u.ID = ""
	update := bson.M{
		"$set": u,
	}
	_, err = db.Collection(CollectionName).UpdateByID(ctx, hex, update, nil)
	if err != nil {
		return nil, err
	}
	u.ID = hex.Hex()
	return u, nil
}

func (u *User) Delete(db *mongo.Database) (int64, error) {
	var user User
	if u.ID == "" {
		return 0, errors.New("invalid id")
	}
	hex, err := primitive.ObjectIDFromHex(u.ID)
	if err != nil {
		return 0, err
	}
	query := bson.M{"_id": hex}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor := db.Collection(CollectionName).FindOne(ctx, query)
	err = cursor.Decode(&user)
	if err != nil {
		return 0, err
	}
	if user.ID == "" {
		return 0, errors.New("invalid id")
	}
	result, err := db.Collection(CollectionName).DeleteOne(ctx, query)
	if err != nil {
		return 0, err
	}
	if result.DeletedCount > 0 {
		return result.DeletedCount, nil
	}
	return 0, errors.New("unable to remove user")
}

func (u *User) FindById(db *mongo.Database) (*User, error) {
	var user User
	hex, err := primitive.ObjectIDFromHex(u.ID)
	if err != nil {
		return nil, err
	}
	query := bson.M{"_id": hex}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor := db.Collection(CollectionName).FindOne(ctx, query)
	err = cursor.Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) FindAll(from, to int64, db *mongo.Database) ([]*User, error) {
	var users []*User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	query := bson.D{}
	options := options.Find()
	options.SetSort(map[string]int{"createdAt": -1})
	options.SetSkip(from)
	options.SetLimit(to)
	cursor, err := db.Collection(CollectionName).Find(ctx, query, options)
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var u User
		err = cursor.Decode(&u)
		if err != nil {
			return nil, err
		}
		users = append(users, &u)
	}
	return users, nil
}
