package products31

import (
	"context"
	"github.com/google/uuid"
	config "github.com/kirigaikabuto/common-lib31"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
)

type productStore struct {
	collection *mongo.Collection
}

func NewProductStore(config config.MongoConfig) (ProductStore, error) {
	clientOptions := options.Client().ApplyURI("mongodb://" + config.Host + ":" + config.Port)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	db := client.Database(config.Database)
	err = db.CreateCollection(context.TODO(), config.CollectionName)
	if err != nil && !strings.Contains(err.Error(), "NamespaceExists") {
		return nil, err
	}
	collection := db.Collection(config.CollectionName)
	return &productStore{collection: collection}, nil
}

func (p *productStore) Create(product *Product) (*Product, error) {
	token := uuid.New()
	product.Id = token.String()
	_, err := p.collection.InsertOne(context.TODO(), product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *productStore) GetById(id string) (*Product, error) {
	filter := bson.D{{"id", id}}
	product := &Product{}
	err := p.collection.FindOne(context.TODO(), filter).Decode(&product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *productStore) List() ([]Product, error) {
	filter := bson.D{}
	products := []Product{}
	cursor, err := p.collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.TODO(), &products)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p *productStore) Delete(id string) error {
	return nil
}
