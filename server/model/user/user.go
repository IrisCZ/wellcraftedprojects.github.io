package user

import (
  "github.com/jjballano/wellcraftedprojects/database"
  "github.com/jjballano/wellcraftedprojects/crypto"
  "gopkg.in/mgo.v2/bson"
  "net/http"
  "io/ioutil"
  "encoding/json"
  "errors"
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

func (user *User) UnmarshalHTTP(request *http.Request) error {
  defer request.Body.Close()
  bodySave, _ := ioutil.ReadAll(request.Body)
  error := json.Unmarshal(bodySave, user)
  if error != nil{
    return error
  }
  if len(user.Email) < 1 || len(user.Password) < 1 {
    return errors.New("Email and password are mandatory")
  }
  return nil
}