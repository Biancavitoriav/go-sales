package repository

import (
	"context"
	"time"
	"vendas/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerRepository interface {
	Save(customer *models.Customer) error
	FindAll() ([]*models.Customer, error)
	FindByID(id string) (*models.Customer, error)
}

type customerRepository struct {
	db *mongo.Database
}

func NewCustomerRepository(db *mongo.Database) CustomerRepository {
	return &customerRepository{db: db}
}

type customerDocument struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `bson:"name"`
	Phone string             `bson:"phone"`
}

func (r *customerRepository) Save(customer *models.Customer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := r.db.Collection("customers")

	doc := customerDocument{
		Name:  customer.Name,
		Phone: customer.Phone,
	}

	result, err := collection.InsertOne(ctx, doc)
	if err != nil {
		return err
	}

	customer.ID = result.InsertedID.(primitive.ObjectID).Hex()
	return nil
}

func (r *customerRepository) FindAll() ([]*models.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := r.db.Collection("customers")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var customers []*models.Customer
	for cursor.Next(ctx) {
		var doc customerDocument
		if err := cursor.Decode(&doc); err != nil {
			return nil, err
		}

		customers = append(customers, &models.Customer{
			ID:    doc.ID.Hex(),
			Name:  doc.Name,
			Phone: doc.Phone,
		})
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return customers, nil
}

func (r *customerRepository) FindByID(id string) (*models.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := r.db.Collection("customers")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var doc customerDocument
	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&doc)
	if err != nil {
		return nil, err
	}

	return &models.Customer{
		ID:    doc.ID.Hex(),
		Name:  doc.Name,
		Phone: doc.Phone,
	}, nil
}
