package db

import (
	"context"
	"os"

	"github.com/by-sabbir/gh-webhook-consumer/internal/hooks"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// func convertJson(gi *hooks.GitIssues) (GitIssues, error) {
// 	var bsonData GitIssues
// 	jsonByte, err := json.Marshal(gi)
// 	if err != nil {
// 		return GitIssues{}, err
// 	}
// 	bson.Unmarshal(jsonByte, bsonData)

// 	return bsonData, nil
// }

func (d *DataBase) GetIssues(ctx context.Context, gi hooks.GitIssues) (string, error) {
	coll := d.Client.Database(os.Getenv("DB_NAME")).Collection("issues")
	// issuesBson, err := convertJson(&gi)
	// if err != nil {
	// 	return "", err
	// }

	res, err := coll.InsertOne(ctx, gi[0])
	if err != nil {
		logrus.Error("insert failed: ", err)
		return "", err
	}

	logrus.Info("inserted successfully: ", gi)
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}
