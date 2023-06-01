package models

type Category struct {
	ID          string `bson:"_id,omitempty" json:"id"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
}
