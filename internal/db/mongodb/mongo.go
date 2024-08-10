package mongodb

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/zeze322/wt-guided-weaponry/models"
)

type Store interface {
	Weapons(context.Context) ([]*models.Params, error)
	WeaponByName(context.Context, string) (*models.Params, error)
	WeaponsByCategory(context.Context, string) ([]*models.Params, error)
	InsertWeapon(context.Context, *models.Params) error
	UpdateWeapon(context.Context, string, *models.Params) error
	DeleteWeapon(context.Context, string) error
	SearchWeapon(context.Context, string) ([]*models.Params, error)
}

type MongoClient struct {
	client          *mongo.Client
	mongoURL        string
	mongoDatabase   string
	mongoCollection string
}

func New(ctx context.Context, mongoURL, mongoDatabase, mongoCollection string) (*MongoClient, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return &MongoClient{
		client:          client,
		mongoURL:        mongoURL,
		mongoDatabase:   mongoDatabase,
		mongoCollection: mongoCollection,
	}, nil
}

func (m *MongoClient) Close(ctx context.Context) error {
	return m.client.Disconnect(ctx)
}

func (m *MongoClient) Weapons(ctx context.Context) ([]*models.Params, error) {
	coll := m.client.Database(m.mongoDatabase).Collection(m.mongoCollection)

	filter := bson.M{}

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
		return nil, fmt.Errorf("weapons not found")
	}

	return weapons, nil
}

func (m *MongoClient) WeaponByName(ctx context.Context, name string) (*models.Params, error) {
	coll := m.client.Database(m.mongoDatabase).Collection(m.mongoCollection)

	weapon := new(models.Params)

	filter := bson.M{"name": name}

	err := coll.FindOne(ctx, filter).Decode(weapon)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, fmt.Errorf("weapon not found: %s", name)
	} else if err != nil {
		return nil, fmt.Errorf("failed to get weapon: %s", name)
	}

	return weapon, nil
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
		return nil, fmt.Errorf("weapons not found for %s category", category)
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
		return fmt.Errorf("weapon already exists: %s", params.Name)
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
		return fmt.Errorf("weapon not found: %s", name)
	}

	return nil
}

func (m *MongoClient) DeleteWeapon(ctx context.Context, name string) error {
	coll := m.client.Database(m.mongoDatabase).Collection(m.mongoCollection)

	filter := bson.M{"name": name}

	res, err := coll.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return fmt.Errorf("weapon not found: %s", name)
	}

	return nil
}

func (m *MongoClient) SearchWeapon(ctx context.Context, keyWord string) ([]*models.Params, error) {
	coll := m.client.Database(m.mongoDatabase).Collection(m.mongoCollection)

	filter := bson.D{{"$text", bson.D{{"$search", keyWord}}}}

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
		return nil, fmt.Errorf("nothing found: %s", keyWord)
	}

	return weapons, nil
}
