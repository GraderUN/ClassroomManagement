package main

import (
	"context"
	"github.com/urfave/cli/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var collection *mongo.Collection
var ctx = context.TODO()

func init() {
	clientOptions := options.Client().ApplyURI(
		"mongodb+srv://Sanhernandezmon:Shm1qazz@realmcluster.xofvc.mongodb.net/GraderDB?retryWrites=true&w=majority",
	)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database("GraderDB").Collection("Courses")

}

func main() {
	app := &cli.App{
		Name:     "GraderCoursesManagement",
		Usage:    "Ms for manage the courses and classrooms asignements on GraderApp",
		Commands: []*cli.Command{},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
