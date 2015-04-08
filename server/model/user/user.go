package user

import (
  "github.com/IrisCZ/wellcraftedprojects/database"
  "github.com/IrisCZ/wellcraftedprojects/crypto"
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

type Credentials struct {
  Login string
  Password string
}

var db database.Database

func Init(database database.Database){
  db = database
}

func (user *User) Save() (string,error){
  user.Password = crypto.EncodePassword(user.Password)
  if FindBy(user.Email) != nil {
    return "",errors.New("user.exists.error")
  }
  newUser,_ := db.Save(user, collectionName)
  return newUser,nil
}

func (user *User) SetId(id bson.ObjectId){
  user.Id = id
}

func FindBy(email string) *User {
  params := make(map[string]string)
  params["email"] = email
  user := new(User)
  error := db.FindOne(collectionName, params, user)
  if error != nil {
    return nil
  }
  return user
}

func Login(email string, password string) *User {
  passwordEncoded := crypto.EncodePassword(password)
  user := FindBy(email)
  if user != nil && user.Password == passwordEncoded {
    return user
  }
  return nil
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

func (credentials *Credentials) UnmarshalHTTP(request *http.Request) error {
  defer request.Body.Close()
  bodySave, _ := ioutil.ReadAll(request.Body)
  error := json.Unmarshal(bodySave, credentials)
  if error != nil{
    return error
  }
  if len(credentials.Login) < 1 || len(credentials.Password) < 1 {
    return errors.New("Login and password are mandatory")
  }
  return nil
}