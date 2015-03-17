package http

import (
    "fmt"
    "github.com/gorilla/mux"
    "github.com/jjballano/wellcraftedprojects/model/user"
    "net/http"
    "os"
)

type ResponseParam struct {
    Name string
    Value string
}

func newUser(response http.ResponseWriter, request *http.Request) {
  defer request.Body.Close()
  theUser := user.User{}
  ParseBodyTo(request.Body, &theUser)
  id := theUser.Save()

  params := make([]ResponseParam,1)
  params[0] = ResponseParam{Name:"id", Value:id}
  ParseResponseTo(response, "OK", params)
}

func notFound(response http.ResponseWriter, request *http.Request) {
  fmt.Fprintf(response, "Error")
}

func StartApi() {
  router := mux.NewRouter()
  router.HandleFunc("/user/new", newUser).Methods("POST")
  router.HandleFunc("/", notFound)
  http.Handle("/", router)
  port := os.Getenv("PORT")
  if len(port) < 1 {
    port = "1337"
  }
  http.ListenAndServe(":"+port, nil)
}