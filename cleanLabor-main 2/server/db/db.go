package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"my-app/model"
	"net/http"
)

type DB interface {
	GetTechnologies() ([]*model.Technology, error)
	GetTaskers() ([]*model.Tasker, error)
	AddTasker(tasker model.Tasker) ([]*model.Tasker, error)
}

type MongoDB struct {
	collection *mongo.Collection
}

func NewMongo(client *mongo.Client) DB {
	tech := client.Database("tech").Collection("tech")
	return MongoDB{collection: tech}
}
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
func (m MongoDB) GetTechnologies() ([]*model.Technology, error) {

	res, err := m.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Println("Error while fetching technologies:", err.Error())
		return nil, err
	}
	var tech []*model.Technology
	err = res.All(context.TODO(), &tech)
	if err != nil {
		log.Println("Error while decoding technologies:", err.Error())
		return nil, err
	}
	return tech, nil
}

// Function to display all the Taskers in the Database (For test use) [Yeshan Li]
func (m MongoDB) GetTaskers() ([]*model.Tasker, error) {
	res, err := m.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Println("Error while fetching taskers:", err.Error())
		return nil, err
	}
	var tsk []*model.Tasker
	err = res.All(context.TODO(), &tsk)
	if err != nil {
		log.Println("Error while decoding taskers:", err.Error())
		return nil, err
	}
	return tsk, nil
}

// Function to Add tasker to the Database [Yeshan Li]
func (m MongoDB) AddTasker(tasker model.Tasker) ([]*model.Tasker, error) {
	m.collection.InsertOne(context.TODO(), tasker)
	res, err := m.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Println("Error while fetching taskers:", err.Error())
		return nil, err
	}
	var tsk []*model.Tasker
	err = res.All(context.TODO(), &tsk)
	if err != nil {
		log.Println("Error while decoding taskers:", err.Error())
		return nil, err
	}
	return tsk, nil
}
