package models

type Customer struct {
	ID    string `json:"id" bson:"_id,omitempty"`
	Name  string `json:"name" bson:"name"`
	Phone string `json:"phone" bson:"phone"`
}
