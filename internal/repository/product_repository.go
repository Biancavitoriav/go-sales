package repository

import (
	"context"
	"time"
	"vendas/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository interface {
	Save(product *models.Product) error
	FindAll() ([]*models.Product, error)
	FindByID(id string) (*models.Product, error)
}

type productRepository struct {
	db *mongo.Database
}

func NewProductRepository(db *mongo.Database) ProductRepository {
	return &productRepository{db: db}
}

type productDocument struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `bson:"name"`
	Price float64            `bson:"price"`
}

func (r *productRepository) Save(product *models.Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := r.db.Collection("products")

	doc := productDocument{
		Name:  product.Name,
		Price: product.Price,
	}

	result, err := collection.InsertOne(ctx, doc)
	if err != nil {
		return err
	}

	product.ID = result.InsertedID.(primitive.ObjectID).Hex()
	return nil
}

func (r *productRepository) FindAll() ([]*models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := r.db.Collection("products")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []*models.Product
	for cursor.Next(ctx) {
		var doc productDocument
		if err := cursor.Decode(&doc); err != nil {
			return nil, err
		}

		products = append(products, &models.Product{
			ID:    doc.ID.Hex(),
			Name:  doc.Name,
			Price: doc.Price,
		})
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *productRepository) FindByID(id string) (*models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := r.db.Collection("products")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var doc productDocument
	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&doc)
	if err != nil {
		return nil, err
	}

	return &models.Product{
		ID:    doc.ID.Hex(),
		Name:  doc.Name,
		Price: doc.Price,
	}, nil
}
