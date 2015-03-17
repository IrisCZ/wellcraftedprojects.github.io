package mongo

import (
    "gopkg.in/mgo.v2"
    "github.com/jjballano/wellcraftedprojects/model"
)

var database string
var host string

type Mongo struct{}

func Init(aHost string, aDatabase string){
  host = aHost
  database = aDatabase
}

func (mongo Mongo) Save(obj model.Model, collectionName string) string {
    session, err := mgo.Dial(host)
    if err != nil{
        panic(err)
    }
    defer session.Close()

    database := session.DB(database)
    collection := database.C(collectionName)
    err = collection.Insert(obj)
    if err != nil{
        return ""
    }
    return "IDDD"
}