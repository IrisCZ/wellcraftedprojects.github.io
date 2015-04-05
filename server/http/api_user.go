package http

import (
  "net/http"
  "github.com/IrisCZ/wellcraftedprojects/model/user"
  "github.com/IrisCZ/wellcraftedprojects/crypto"
)

func NewUser(response http.ResponseWriter, request *http.Request) {
  theUser,error := parseUser(request)
  if error != nil {
    returnError(response, error)
    return
  }
//  params := make([]ResponseParam,1)

  id,error := theUser.Save()
  if error != nil {
    returnError(response, error)
  } else {
    parseResponseTo(response, "OK", map[string]interface{}{"id":id})
  }
}

func Login(response http.ResponseWriter, request *http.Request) {
  credentials, error := parseCredentials(request)
  if error != nil {
    returnError(response, error)
    return
  }
  user := user.Login(credentials.Login,credentials.Password)
  if user != nil {
    userData := map[string]interface{}{"login":user.Email}
    parseResponseTo(response, "OK",map[string]interface{}{"token":crypto.CreateToken(userData)})
  } else {
    returnError(response,nil)
  }
}

func parseUser(request *http.Request)(*user.User, error){
  theUser := new(user.User)
  return theUser, getEntity(request, theUser)
}

func parseCredentials(request *http.Request)(*user.Credentials, error){
    credentials := new(user.Credentials)
    return credentials, getEntity(request, credentials)
}

func returnError(response http.ResponseWriter, error error){
  parseResponseTo(response, "ERROR", map[string]interface{}{"error":error.Error()})
}