package model

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
type Model struct {
	client *mongo.Client
	collectionPersons *mongo.Collection
}


type Person struct{
	Name string `bson:string`
	Age string  `bson:string`
	Pnum string  `bson:string`
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
		db := r.client.Database("go-ready")
		r.collectionPersons = db.Collection("tPerson")
	}
  return r, nil
}

func (m *Model) GetPersonByName(name string) (Person, error){
   filter := bson.M{"name":name}
   var result Person
   err := m.collectionPersons.FindOne(context.TODO(), filter).Decode(&result)
   if err == mongo.ErrNoDocuments {
	   fmt.Printf("No document was found with the name %s\n", name)
	   return Person{}, errors.New("error")
   } else if err != nil {
	   panic(err)
   }

  return result, nil
  
}


func (m *Model) GetPersonByPnum(Pnum string) (Person, error){
	filter := bson.M{"pnum":Pnum}
	var result Person
	err := m.collectionPersons.FindOne(context.TODO(), filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the name %s\n", Pnum)
		return Person{}, errors.New("error")
	} else if err != nil {
		panic(err)
	}
 
   return result, nil
   
 }

 func (m *Model) JoinPerson(personInfo Person) (any,error){
	 info := bson.M{"name":personInfo.Name, "age":personInfo.Age, "pnum":personInfo.Pnum}
	 result, err := m.collectionPersons.InsertOne(context.TODO(), info)
	 if err != nil {
		 return "",errors.New("error")
	 }
	 return result.InsertedID, nil
 }

 func (m *Model) UpdatePerson(personInfo Person) (any, error){
	 filter := bson.M{"name":personInfo.Name}
	 update := bson.D{{"$set", bson.D{{"age",personInfo.Age}}}}
	 result, err := m.collectionPersons.UpdateOne(context.TODO(), filter, update)
	  if err != nil {
		panic(err)
	}
	return result.ModifiedCount, nil
	
 }



 func (m *Model) DeletePerson(personInfo Person) (any, error){
	filter := bson.M{"name":personInfo.Name}

	result, err := m.collectionPersons.DeleteOne(context.TODO(), filter)
	 if err != nil {
	   panic(err)
   }
   return result.DeletedCount, nil
   
}