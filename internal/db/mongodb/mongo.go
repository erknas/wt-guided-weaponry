package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/zeze322/wt-guided-weaponry/models"
)

type Store interface {
	Categories(context.Context) ([]models.Category, error)
	Weapons(context.Context) ([]*models.Params, error)
	WeaponsByCategory(context.Context, string) ([]*models.Params, error)
	InsertWeapon(context.Context, *models.Params) error
	UpdateWeapon(context.Context, string, *models.Params) error
	SearchWeapon(context.Context, string) ([]models.Name, error)
}

type MongoClient struct {
	client          *mongo.Client
	mongoDatabase   string
	mongoCollection string
}

func New(ctx context.Context, mongoURI, mongoDatabase, mongoCollection string) (*MongoClient, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return &MongoClient{
		client:          client,
		mongoDatabase:   mongoDatabase,
		mongoCollection: mongoCollection,
	}, nil
}

func (m *MongoClient) Close(ctx context.Context) error {
	return m.client.Disconnect(ctx)
}

func (m *MongoClient) Categories(ctx context.Context) ([]models.Category, error) {
	coll := m.client.Database(m.mongoDatabase).Collection(m.mongoCollection)

	filter := bson.M{"categories": "exists"}

	cursor, err := coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var categories []models.Category

	if err := cursor.All(ctx, &categories); err != nil {
		return nil, err
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

func (m *MongoClient) Weapons(ctx context.Context) ([]*models.Params, error) {
	coll := m.client.Database(m.mongoDatabase).Collection(m.mongoCollection)

	filter := bson.M{"name": bson.M{"$ne": nil}}

	cursor, err := coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var weapons []*models.Params

	if err := cursor.All(ctx, &weapons); err != nil {
		return nil, err
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return weapons, nil
}

func (m *MongoClient) WeaponsByCategory(ctx context.Context, category string) ([]*models.Params, error) {
	coll := m.client.Database(m.mongoDatabase).Collection(m.mongoCollection)

	filter := bson.M{"category": category}

	cursor, err := coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var weapons []*models.Params

	if err := cursor.All(ctx, &weapons); err != nil {
		return nil, err
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(weapons) == 0 {
		return nil, fmt.Errorf("")
	}

	return weapons, nil
}

func (m *MongoClient) InsertWeapon(ctx context.Context, params *models.Params) error {
	coll := m.client.Database(m.mongoDatabase).Collection(m.mongoCollection)

	weapon := models.NewWeapon(params)

	filter := bson.M{"name": params.Name}
	count, err := coll.CountDocuments(ctx, filter)
	if err != nil {
		return err
	}

	if count != 0 {
		return fmt.Errorf("")
	}

	_, err = coll.InsertOne(ctx, weapon)
	if err != nil {
		return err
	}

	return nil
}

func (m *MongoClient) UpdateWeapon(ctx context.Context, name string, params *models.Params) error {
	coll := m.client.Database(m.mongoDatabase).Collection(m.mongoCollection)

	update := bson.M{"$set": models.UpdateWeaponParams(params)}
	filter := bson.M{"name": name}

	res, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if res.MatchedCount == 0 {
		return fmt.Errorf("")
	}

	return nil
}

func (m *MongoClient) SearchWeapon(ctx context.Context, keyWord string) ([]models.Name, error) {
	coll := m.client.Database(m.mongoDatabase).Collection(m.mongoCollection)

	filter := bson.D{{Key: "name", Value: primitive.Regex{Pattern: keyWord, Options: "i"}}}

	cursor, err := coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var weapons []models.Name

	if err := cursor.All(ctx, &weapons); err != nil {
		return nil, err
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return weapons, nil
}
