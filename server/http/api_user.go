package http

import (
  "net/http"
  "github.com/jjballano/wellcraftedprojects/model/user"
  "github.com/jjballano/wellcraftedprojects/crypto"
)

func NewUser(response http.ResponseWriter, request *http.Request) {
  theUser,error := parseUser(request)
  if error != nil {
    returnError(response, error)
    return
  }
  id := theUser.Save()

  params := make([]ResponseParam,1)
  params[0] = ResponseParam{Name:"id", Value:id}
  parseResponseTo(response, "OK", params)
}

func Login(response http.ResponseWriter, request *http.Request) {
  login, error := parseLogin(request)
  if error != nil {
    returnError(response, error)
    return
  }
  user := user.FindBy(login.Login,login.Password)
  if user != nil {
    userData := map[string]interface{}{"login":user.Email}
    params := make([]ResponseParam,1)
    params[0] = ResponseParam{Name:"token", Value:crypto.CreateToken(userData)}
    parseResponseTo(response, "OK",params)
  } else {
    parseResponseTo(response, "ERROR",nil)
  }
}

func parseUser(request *http.Request)(*user.User, error){
  theUser := new(user.User)
  return theUser, getEntity(request, theUser)
}

func parseLogin(request *http.Request)(*user.Login, error){
    login := new(user.Login)
    return login, getEntity(request, login)
}

func returnError(response http.ResponseWriter, error error){
  params := make([]ResponseParam,1)
  params[0] = ResponseParam{Name:"error", Value:error.Error()}
  parseResponseTo(response, "ERROR", params)
}