package model

type IV_model struct{
	Name string  `json:"name" bson:"name"`
	Quantity float32  `json:"quantity" bson:"quantity"`
} 