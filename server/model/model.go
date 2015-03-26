package model
import "gopkg.in/mgo.v2/bson"

type Model interface{
    SetId(id bson.ObjectId)
}