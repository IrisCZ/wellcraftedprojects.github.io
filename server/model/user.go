package model

import (
  "github.com/jjballano/wellcraftedprojects/database"
  "github.com/jjballano/wellcraftedprojects/crypto"
)

const collectionName string = "users"


type User struct {
  Id string
  Email string
  Password string
}

var mongo = database.Mongo{}

func (user *User) Save() string{
  user.Password = crypto.EncodePassword(user.Password)
    return mongo.Save(user, collectionName)
}
