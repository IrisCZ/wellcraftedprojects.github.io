package user

import (
  "github.com/jjballano/wellcraftedprojects/database"
  "github.com/jjballano/wellcraftedprojects/crypto"
    "gopkg.in/mgo.v2/bson"
)

const collectionName string = "users"


type User struct {
  Id bson.ObjectId `_id,omitempty`
  Email string
  Password string
}

var db database.Database

func Init(database database.Database){
  db = database
}

func (user *User) Save() string{
  user.Password = crypto.EncodePassword(user.Password)
  newUser,_ := db.Save(user, collectionName)
  return newUser
}

func (user *User) SetId(id bson.ObjectId){
  user.Id = id
}