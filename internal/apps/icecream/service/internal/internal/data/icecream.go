package data

import "go.mongodb.org/mongo-driver/bson/primitive"

type IceCreamField string

type IceCream struct {
	Id                    primitive.ObjectID `bson:"_id"`
	ProductId             string             `bson:"product_id"`
	Name                  string             `bson:"name"`
	ImageClosed           string             `bson:"image_closed"`
	ImageOpen             string             `bson:"image_open"`
	Description           string             `bson:"description"`
	Story                 string             `bson:"story"`
	SourcingValues        []string           `bson:"sourcing_values"`
	Ingredients           []string           `bson:"ingredients"`
	AllergyInfo           string             `bson:"allergy_info"`
	DietaryCertifications string             `bson:"dietary_certifications"`
}

type IceCreamListResult struct {
	IceCreams []*IceCream
	LastId    string
}

type IceCreamSearchOptions struct {
	LastId string
	Limit  int64
}
