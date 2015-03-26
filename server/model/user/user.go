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

type Login struct {
  Login string
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

func FindBy(email string, password string) *User {
  params := make(map[string]string)
  params["email"] = email
  params["password"] = crypto.EncodePassword(password)
  user := new(User)
  error := db.FindOne(collectionName, params, user)
  if error != nil {
    return nil
  }
  return user
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

func (login *Login) UnmarshalHTTP(request *http.Request) error {
    defer request.Body.Close()
    bodySave, _ := ioutil.ReadAll(request.Body)
    error := json.Unmarshal(bodySave, login)
    if error != nil{
        return error
    }
    if len(login.Login) < 1 || len(login.Password) < 1 {
        return errors.New("Login and password are mandatory")
    }
    return nil
}