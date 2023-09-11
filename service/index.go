package service

import (
	"context"
	"log"
	"prac_proj/interfaces"
	"prac_proj/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type Invent struct {
	ctx             context.Context
	mongoCollection *mongo.Collection
}

// Create__IV implements interfaces.IV_int.
func (q *Invent) Create__IV(value *model.IV_model) (string, error) { 

	_,err:=q.mongoCollection.InsertOne(q.ctx,value)
	if err!=nil{
		return "",err
	}
	return "success",nil
}

func InitInventory(collection *mongo.Collection, ctx context.Context) interfaces.IV_int {


	return &Invent{ctx, collection}
}
