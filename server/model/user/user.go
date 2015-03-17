package user

import (
  "github.com/jjballano/wellcraftedprojects/database"
  "github.com/jjballano/wellcraftedprojects/crypto"
  "io"
  "io/ioutil"
  "encoding/json"
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

func From(body io.ReadCloser) User {
  aUser := User{}
  bodySave, _ := ioutil.ReadAll(body)
  json.Unmarshal(bodySave, &aUser)
  return aUser
}
