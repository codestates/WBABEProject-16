package model

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
type Model struct {
	client *mongo.Client
	colPizzaCategory *mongo.Collection
	colOrderInfo *mongo.Collection
	colPersonInfo *mongo.Collection
}

type OrderInfo struct {
  ID      primitive.ObjectID  `json:"_id" bson:"_id,omitempty"`
  PizzaId string `json:"pizza_id"`
  Size string `json:"size"`
  Amount int `json:"amount"`
  personId string `json:"person_id"`
  Status string `json:"status"`
  created_at time.Time `json:"created_at"`
  updated_at time.Time `json:"updated_at"`
}

type OrderPersonInfo struct {
	ID      primitive.ObjectID  `json:"_id" bson:"_id,omitempty"`
	Name string `json:"name"`
	Phone string `json:"phone"`
	Address string `json:"address"`
	PizzaId string `json:"pizza_id:`
	createdAt time.Time `json:"created_at"`
}
type PizzaCategory struct {
	ID      primitive.ObjectID  `json:"_id" bson:"_id,omitempty"`
	Name string `json:name binding:"required"`
	Des string `json:des binding:"required"`
	M int `json:m binding:"required"`
	L int `json:l binding:"required"`
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
		r.colPersonInfo = db.Collection("person_info")
		r.colOrderInfo = db.Collection("order_order");
	}
  return r, nil
}


func (m *Model) AddCategory(category PizzaCategory) (bool,error) {

	_, findErr := m.findByName(category.Name)
	if findErr == nil {
		return false, errors.New("error")
	}

	doc := bson.M{
		"name":category.Name,
		"des":category.Des,
		"msizeprice":category.M,
		"lsizeprice":category.L,
		"order_status":category.Order_status,
		"limit_order": category.Limit_Order,
		"updated_at" : time.Now(),
		"deleted_at" : nil,
	}
   result, err := m.colPizzaCategory.InsertOne(context.TODO(), doc)
   if err != nil {
	return false, errors.New("error")
   }
var stringObjectId string = result.InsertedID.(primitive.ObjectID).String()
 fmt.Println(stringObjectId)
	return true,nil
}


func (m *Model) UpdateCategory(category PizzaCategory) (bool,error) {
	
	_, findErr := m.findByName(category.Name)
	if findErr != nil {
		return false, errors.New("error")
	}
	
    filter := bson.M{"name":category.Name}
	update := bson.M{
		"$set": bson.M{
		"des":category.Des,
		"msizeprice":category.M,
		"lsizeprice":category.L,
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


func (m *Model) DeleteByName(name string) (bool, error){
	findResult, findErr := m.findByName(name)
	if findErr != nil {
		return false, errors.New("error")
	}
	filter := bson.M{"name":name}
	update := bson.M{
		"$set": bson.M{
		"des":findResult.Des,
		"msizeprice":findResult.M,
		"lsizeprice":findResult.L,
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

func (m *Model) getCategory() {
	
}

func (m *Model) OrderPizza(orderInfo OrderInfo,orderPersonInfo OrderPersonInfo) (string, error) {
	type OrderInfo struct {
		PizzaId primitive.ObjectID `json:"pizza_id"`
		Size string `json:"size"`
		Amount int `json:"amount"`
		personId primitive.ObjectID `json:"person_id"`
		Status string `json:"status"`
		created_at time.Time `json:"created_at"`
		updated_at time.Time `json:"updated_at"`
	  }
	  
	  type OrderPersonInfo struct {
		  Name string `json:"name"`
		  Phone string `json:"phone"`
		  Address string `json:"address"`
		  PizzaId primitive.ObjectID `json:"pizza_id:`
		  createdAt time.Time `json:"created_at"`
	  }

	  pizzaId, err := primitive.ObjectIDFromHex(orderInfo.PizzaId)
		if err != nil {
		  panic(err)
		}
      m.findByName("string")
	  m.findCategoryById(pizzaId)
	 orderPersonInfoDoc := bson.M{
		 "name": orderPersonInfo.Name,
		 "phone" : orderPersonInfo.Phone,
		 "address" : orderPersonInfo.Address,
		 "pizza_id": pizzaId,
		 "create_at" : time.Now(),
		}
		personResult, personErr := m.colPersonInfo.InsertOne(context.TODO(), orderPersonInfoDoc)
		if personErr != nil {
		 return "", errors.New("error")
		}
		//var orderPeronId string = personResult.InsertedID.(primitive.ObjectID).String()
	
	 orderInfoDocs := bson.M{
		 "pizza_id":pizzaId,
		 "size" : orderInfo.Size,
		 "amount" : orderInfo.Amount,
		 "person_id":personResult.InsertedID,
		 "status":orderInfo.Status,
		 "created_at" : time.Now(),
		 "updated_at" : nil,
	 } 	
	 orderResult, orderErr := m.colOrderInfo.InsertOne(context.TODO(),orderInfoDocs)
	 if orderErr != nil {
		return "", errors.New("error")
	   }
	var orderId string = orderResult.InsertedID.(primitive.ObjectID).String()
	   return orderId,nil
}

func(m *Model) UpdateOrderStatus(objId string, orderStatus string)(int64,error){
	orderId, err := primitive.ObjectIDFromHex(objId)
		if err != nil {
		  panic(err)
		  return 0, errors.New("nil이아님")
		}

	_, findErr := m.findOrderById(orderId)	
	if findErr != nil {
		return 0, errors.New("orderId not found")
	}
	filter := bson.M{"_id":orderId}
	update := bson.M{"$set": bson.M{"order_status":orderStatus}}
	result, err := m.colOrderInfo.UpdateOne(context.TODO(),filter,update)
	  if err != nil {
		  return 0, errors.New("nil이아님")
	  }
	return  result.ModifiedCount, nil
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
	fmt.Println(result)
	return result, nil
}

func (m *Model) findCategoryById(id primitive.ObjectID)(PizzaCategory,error){
	filter := bson.M{"_id":id}
	var result PizzaCategory
	err := m.colPizzaCategory.FindOne(context.TODO(), filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the id %s\n", id)
		return PizzaCategory{}, errors.New("error")
	} else if err != nil {
		panic(err)
	}	
	fmt.Println(result)
	return result, nil
}

func (m *Model) findOrderById(id primitive.ObjectID)(OrderInfo,error){
	filter := bson.M{"_id":id}
	var result OrderInfo
	err := m.colOrderInfo.FindOne(context.TODO(), filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the id %s\n", id)
		return OrderInfo{}, errors.New("error")
	} else if err != nil {
		panic(err)
	}	
	return result, nil
}


