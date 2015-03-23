package http

import (
  "net/http"
  "github.com/jjballano/wellcraftedprojects/model/user"
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
  _,error := parseUser(request)
  if error != nil {
    returnError(response, error)
    return
  }
}

func parseUser(request *http.Request)(*user.User, error){
  theUser := new(user.User)
  error := getEntity(request, theUser)
  return theUser, error
}

func returnError(response http.ResponseWriter, error error){
  params := make([]ResponseParam,1)
  params[0] = ResponseParam{Name:"error", Value:error.Error()}
  parseResponseTo(response, "ERROR", params)
}