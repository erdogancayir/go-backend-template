package bootstrap

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDatabase(env *Env) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbHost := env.DBHost
	dbPort := env.DBPort
	dbUser := env.DBUser
	dbPass := env.DBPass

	/* log.Fatal("->>  ", dbUser, " , ", dbPass, " , ", dbHost, ",", dbPort) */
	mongodbURI := fmt.Sprintf("mongodb://%s:%s@%s:%s", dbUser, dbPass, dbHost, dbPort)
	if dbUser == "" || dbPass == "" {
		mongodbURI = fmt.Sprintf("mongodb://%s:%s", dbHost, dbPort)
	}

	//	Bu satırda bir MongoDB istemci (client) oluşturuluyor.
	//	ApplyURI kodu ile MongoDB sunucusunun adresi (URI) veriliyor.
	client, err := mongo.NewClient(options.Client().ApplyURI(mongodbURI))
	if err != nil {
		return nil, err
	}

	//	Önceden oluşturulan istemci ile MongoDB sunucusuna bağlantı kuruluyor.
	//	Burada ctx olarak tanımladığımız context kullanılıyor.
	//	Yani bağlantı işlemi 10 saniye içinde tamamlanmazsa otomatik olarak iptal edilir.
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	//  Bağlantının başarılı olup olmadığını kontrol etmek için,
	//	MongoDB sunucusuna bir "ping" işlemi gönderiliyor.
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func CloseMongoDBConnection(client *mongo.Client) error {
	if client == nil {
		return nil
	}

	err := client.Disconnect(context.TODO())
	if err != nil {
		return err
	}

	log.Println("Connection to MongoDB closed.")
	return nil
}
