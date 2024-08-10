package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m *MongoClient) CreateIndex(ctx context.Context) error {
	coll := m.client.Database(m.mongoDatabase).Collection(m.mongoCollection)

	model := mongo.IndexModel{Keys: bson.D{{"name", "text"}}}

	_, err := coll.Indexes().CreateOne(ctx, model)
	if err != nil {
		return err
	}

	return nil
}
