package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type ProductMongo struct {
	debug bool
}

func NewProductMongoRepo() *ProductMongo {
	return &ProductMongo{}
}

func (r *ProductMongo) GetMissingMany(client *mongo.Client, offset int64) (*mongo.Cursor, error) {
	pageOptions := options.Find()
	pageOptions.SetSkip(offset)
	pageOptions.SetLimit(100)
	pageOptions.SetProjection(bson.D{{"product_code", 1}})
	filter := bson.D{{"created_at", nil}}
	cur, err := client.Database("viator").Collection("product").Find(context.TODO(), filter, pageOptions)
	return cur, err
}

func (r *ProductMongo) Create(client *mongo.Client, productCode string, ob map[string]interface{}, collection string) error {
	ob["product_code"] = productCode
	updateInfo := bson.M{}
	for k, v := range ob {
		updateInfo[k] = v
	}
	opts := options.Update().SetUpsert(true)
	filter := bson.D{primitive.E{Key: "product_code", Value: productCode}}
	updater := bson.D{primitive.E{Key: "$set", Value: updateInfo}}
	if _, err := client.Database("viator").Collection(collection).UpdateOne(context.Background(), filter, updater, opts); err != nil {
		log.Println(err)
		return err
	}
	return nil

}

func (r *ProductMongo) GetOne(client *mongo.Client, productCode string, collection string) bool {
	filter := bson.D{primitive.E{Key: "product_code", Value: productCode}}
	cur := client.Database("viator").Collection(collection).FindOne(context.TODO(), filter)
	if cur.Err() != nil {
		return false
	}
	t := map[string]interface{}{}
	err := cur.Decode(&t)
	if err != nil {
		return false
	}

	return true

}
