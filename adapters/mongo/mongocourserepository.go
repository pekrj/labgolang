package mongo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/pekrj/labgolang/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoCourseRepository struct {
}

func (mc *MongoCourseRepository) GetAll() ([]domain.Course, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		panic(err)
	}

	collection := client.Database("testing").Collection("numbers")

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	var course domain.Course

	var coursesFromDb []domain.Course = make([]domain.Course, cur.RemainingBatchLength(), cur.RemainingBatchLength())

	for i := 0; cur.Next(ctx); i++ {
		err := cur.Decode(&course)
		if err != nil {
			log.Fatal(err)
		}
		coursesFromDb[i] = course
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	log.Println("Passei até o fim")

	return coursesFromDb, nil
}
