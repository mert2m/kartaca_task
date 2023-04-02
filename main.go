package main

import (
	"context"
	"log"
	"net/http"
	

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	mongoURI = "mongodb://localhost:27017"
)

type Country struct {
	Name string `bson:"name"`
}

func main() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	router.GET("/staj", func(c *gin.Context) {
		ctx := context.Background()
		collection := client.Database("stajdb").Collection("iller")

		// add sample data to the collection if it's empty
		count, err := collection.CountDocuments(ctx, bson.M{})
		if err != nil {
			log.Fatal(err)
		}
		if count == 0 {
			countryDocuments := []interface{}{
				bson.M{"name": "Turkey"},
				bson.M{"name": "United States"},
				bson.M{"name": "France"},
				bson.M{"name": "Japan"},
				bson.M{"name": "Ukraine"},
				bson.M{"name": "Russia"},
				bson.M{"name": "United Kingdom"},
				bson.M{"name": "China"},
				bson.M{"name": "Bosnia Herzegovina"},
				bson.M{"name": "Germany"},
			}
			_, err = collection.InsertMany(ctx, countryDocuments)
			if err != nil {
				log.Fatal(err)
			}
		}

		var country Country
		cur, err := collection.Aggregate(ctx, bson.A{bson.M{"$sample": bson.M{"size": 1}}})
		if err != nil {
			log.Fatal(err)
		}
		defer cur.Close(ctx)
		if cur.Next(ctx) {
			err = cur.Decode(&country)
			if err != nil {
				log.Fatal(err)
			}
		}
		c.JSON(http.StatusOK, gin.H{"name": country.Name})
	})
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Merhaba Python!")
	})
	if err := router.Run(":4444"); err != nil {
		log.Fatal(err)
	}
}
