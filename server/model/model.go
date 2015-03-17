package model
import "gopkg.in/mgo.v2/bson"

type Model interface{
    Save() string
    SetId(id bson.ObjectId)
}