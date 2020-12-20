package global

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//DB hold a database connection
var DB mongo.Database

func connectToDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbUrl))
	if err != nil {
		log.Fatal("Error conncetinhg to DB: ", err.Error())
	}

	//Note that if a function called returns a pointer, then we must also call a pointer and NOT A VALUE
	//Therefore *client
	DB = *client.Database(dbname)
}

//
func NewDBContext(d time.Duration) (context.Context, context.CancelFunc) {

	return context.WithTimeout(context.Background(), d*performance/100)

}

// ConnectToTestDB  overites DB with the test tdatabase
func ConnectToTestDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbUrl))
	if err != nil {
		log.Fatal("Error conncetinhg to DB: ", err.Error())
	}

	//Note that if a function called returns a pointer, then we must also call a pointer and NOT A VALUE
	//Therefore *client
	DB = *client.Database(dbname + "_test")

}
