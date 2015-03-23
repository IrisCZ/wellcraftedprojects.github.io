package http

import (
  "net/http"
  "github.com/jjballano/wellcraftedprojects/model/user"
)

func newUser(response http.ResponseWriter, request *http.Request) {
    theUser := new(user.User)
    error := getEntity(request, theUser)
    params := make([]ResponseParam,1)
    if error != nil {
        params[0] = ResponseParam{Name:"error", Value:error.Error()}
        parseResponseTo(response, "ERROR", params)
        return
    }
    id := theUser.Save()


    params[0] = ResponseParam{Name:"id", Value:id}
    parseResponseTo(response, "OK", params)
}