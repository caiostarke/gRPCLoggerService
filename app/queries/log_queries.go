package queries

import (
	"context"
	"fmt"

	"github.com/caiostarke/apiAndGRPC/logger_service/app/models"
	"github.com/caiostarke/apiAndGRPC/logger_service/logger_pb"
	"go.mongodb.org/mongo-driver/mongo"
)

type LogQueries struct {
	DB *mongo.Client
}

func (q *LogQueries) CreateLog(l *logger_pb.Log) error {
	fmt.Println("Creating log")

	log := models.Log{
		Message: l.Message,
	}

	collection := q.DB.Database("apiAndGRPC").Collection("logs")

	_, err := collection.InsertOne(context.Background(), log)
	if err != nil {
		return err
	}

	return nil
}
