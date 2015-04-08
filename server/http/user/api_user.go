package user

import (
  "net/http"
  "github.com/IrisCZ/wellcraftedprojects/model/user"
  "github.com/IrisCZ/wellcraftedprojects/crypto"
  "github.com/IrisCZ/wellcraftedprojects/http/utils"
)

func New(response http.ResponseWriter, request *http.Request) {
  theUser,error := parseUser(request)
  if error != nil {
    returnError(response, error)
    return
  }

  id,error := theUser.Save()
  if error != nil {
    returnError(response, error)
  } else {
    utils.ParseResponseTo(response, "OK", map[string]interface{}{"id":id})
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
    utils.ParseResponseTo(response, "OK",map[string]interface{}{"token":crypto.CreateToken(userData)})
  } else {
    returnError(response,nil)
  }
}

func parseUser(request *http.Request)(*user.User, error){
  theUser := new(user.User)
  return theUser, utils.GetEntity(request, theUser)
}

func parseCredentials(request *http.Request)(*user.Credentials, error){
  credentials := new(user.Credentials)
  return credentials, utils.GetEntity(request, credentials)
}

func returnError(response http.ResponseWriter, error error){
  utils.ParseResponseTo(response, "ERROR", map[string]interface{}{"error":error.Error()})
}