package db

import (
	"context"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DataBase struct {
	Client *mongo.Client
}

func NewDatabase() (*DataBase, error) {
	connecttionString := fmt.Sprintf(
		"mongodb://%s:%s@%s:%s/?maxPoolSize=%s&w=%s",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_POOL"),
		os.Getenv("WRITE_CONCERN"),
	)

	dbConn, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connecttionString))

	if err != nil {
		return &DataBase{}, fmt.Errorf("db init failed: %w", err)
	}

	return &DataBase{
		Client: dbConn,
	}, nil
}

func (d *DataBase) PingDB() error {
	logrus.Info("Pinging db")
	if err := d.Client.Ping(context.Background(), readpref.Primary()); err != nil {
		logrus.Error("Ping failed: ", err)
		return err
	}
	logrus.Info("Ping Successfull")
	return nil
}
