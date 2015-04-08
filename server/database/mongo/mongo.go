package mongo

import (
    "gopkg.in/mgo.v2"
    "github.com/IrisCZ/wellcraftedprojects/model"
    "gopkg.in/mgo.v2/bson"
    "fmt"
)

var database string
var host string

type Mongo struct{}

func Init(aHost string, aDatabase string){
  host = aHost
  database = aDatabase
}

func (mongo Mongo) Save(obj model.Model, collectionName string) (string, error) {
  session, err := mgo.Dial(host)
  if err != nil{
    panic(err)
  }
  defer session.Close()

  database := session.DB(database)
  collection := database.C(collectionName)
  id := bson.NewObjectId()
  obj.SetId(id)
  err = collection.Insert(obj)
  if err != nil{
    return "", err
  }
  return fmt.Sprintf("%x", string(id)),nil
}

func (mongo *Mongo) FindOne(collectionName string, params map[string]string, obj model.Model) error {
  session, err := mgo.Dial(host)
  if err != nil{
    panic(err)
  }
  defer session.Close()

  database := session.DB(database)
  collection := database.C(collectionName)
  mongoParams := bson.M{}
  for key,value := range params {
    mongoParams[key] = value
  }
  err = collection.Find(mongoParams).One(obj)
  if err != nil{
    return err
  }
  return nil
}

func (mongo *Mongo) FindAll(collectionName string) ([]map[string]interface{}, error) {
  session, err := mgo.Dial(host)
  if err != nil{
    panic(err)
  }
  defer session.Close()

  database := session.DB(database)
  collection := database.C(collectionName)
  var list []map[string]interface{}
  err = collection.Find(nil).All(&list)
  if err != nil{
    return nil,err
  }
  return list,nil
}

