package model

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
type Model struct {
	client *mongo.Client
	colPizzaCategory *mongo.Collection
}


type PizzaCategory struct {
	Name string `json:name binding:"required"`
	Des string `json:des binding:"required"`
	MsizePrice int `json:msizePrice binding:"required"`
	LsizePrice int `json:lsizePrice binding:"required"`
    Order_status bool `json:order_status binding:"required"`
	Limit_Order int `json:limit_order binding:"required"`
	Updated_At time.Time `json:updated_at`
	Deleted_At time.Time `json:deleted_at`
}



func NewModel() (*Model, error){
	r := &Model{}
	var err error
	mgUrl := "mongodb://127.0.0.1:27017"
	if r.client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(mgUrl)); err != nil {
		return nil, err
	}else if err := r.client.Ping(context.Background(), nil); err != nil {
         return nil, err
	} else {
		db := r.client.Database("wba_project")
		r.colPizzaCategory = db.Collection("pizza_category")
	}
  return r, nil
}


func (m *Model) AddCategory(category PizzaCategory) (bool,error) {

    // JData, err := json.Marshal(category)
	// if err != nil {
	// 	return false, errors.New("error")
	// }

	doc := bson.M{
		"name":category.Name,
		"des":category.Des,
		"msizeprice":category.MsizePrice,
		"lsizeprice":category.LsizePrice,
		"order_status":category.Order_status,
		"limit_order": category.Limit_Order,
		"updated_at" : time.Now(),
		"deleted_at" : nil,
	}
   result, err := m.colPizzaCategory.InsertOne(context.TODO(), doc)
   if err != nil {
	return false, errors.New("error")
   }
   fmt.Println(result)
	return true,nil
}


func (m *Model) UpdateCategory(category PizzaCategory) (bool,error) {
	fmt.Print("durl1")
	_, findErr := m.findByName(category.Name)
	if findErr != nil {
		return false, errors.New("error")
	}
	
    filter := bson.M{"name":category.Name}
	update := bson.M{
		"$set": bson.M{
		"des":category.Des,
		"msizeprice":category.MsizePrice,
		"lsizeprice":category.LsizePrice,
		"order_status":category.Order_status,
		"limit_order": category.Limit_Order,
		"updated_at" : time.Now(),
		"deleted_at" : nil,
		},
	}
	result, err := m.colPizzaCategory.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	if result.ModifiedCount == 1 {
		return true, nil
	}
	return false,errors.New("error")
}

func (m *Model) findByName(name string) (PizzaCategory,error){
	filter := bson.M{"name":name}
	
	var result PizzaCategory
	err := m.colPizzaCategory.FindOne(context.TODO(), filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the name %s\n", name)
		return PizzaCategory{}, errors.New("error")
	} else if err != nil {
		fmt.Println(err)
		panic(err)
	}	
	return result, nil
}


func (m *Model) DeleteByName(name string) (bool, error){
	findResult, findErr := m.findByName(name)
	if findErr != nil {
		return false, errors.New("error")
	}
	filter := bson.M{"name":name}
	update := bson.M{
		"$set": bson.M{
		"des":findResult.Des,
		"msizeprice":findResult.MsizePrice,
		"lsizeprice":findResult.LsizePrice,
		"order_status":findResult.Order_status,
		"limit_order": findResult.Limit_Order,
		"updated_at" : findResult.Updated_At,
		"deleted_at" : time.Now(),
		},
	}
	result, err := m.colPizzaCategory.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	if result.ModifiedCount == 1 {
		return true, nil
	}
	return false,errors.New("error")
}