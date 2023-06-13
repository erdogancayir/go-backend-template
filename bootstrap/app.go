package bootstrap

import (
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	Env   *Env
	Mongo *mongo.Client
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	var err error
	//	Environments değerleri ile bağlantı oluşturulur.
	app.Mongo, err = NewMongoDatabase(app.Env)
	if err != nil {
		//	log.Fatalf hata durumunda programın sonlanmasını sağlar.
		// Sonlanmaması için log.Printf
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}
	log.Printf("Connecting to MongoDB created")
	return *app
}

func (app *Application) CloseDBConnection() {
	err := CloseMongoDBConnection(app.Mongo)
	if err != nil {
		log.Fatalf("Error closing MongoDB connection: %v", err)
	}
}
