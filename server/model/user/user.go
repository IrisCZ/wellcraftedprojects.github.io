package user

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

var db database.Database

func Init(database database.Database){
  db = database
}

func (user *User) Save() string{
  user.Password = crypto.EncodePassword(user.Password)
  return db.Save(user, collectionName)
}
