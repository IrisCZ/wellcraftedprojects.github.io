package database


import (
    "gopkg.in/mgo.v2"
)

const databaseName string = "wellcrafted"
const databaseHost string = "localhost"

type Mongo struct{}

func (mongo *Mongo) Save(obj Dao, collectionName string) string {
    session, err := mgo.Dial(databaseHost)
    if err != nil{
        panic(err)
    }
    defer session.Close()

    database := session.DB(databaseName)
    collection := database.C(collectionName)
    err = collection.Insert(obj)
    if err != nil{
        return ""
    }
    return "IDDD"
}