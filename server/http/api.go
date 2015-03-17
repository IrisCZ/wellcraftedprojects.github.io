package http

import (
    "fmt"
    "github.com/gorilla/mux"
    "github.com/jjballano/wellcraftedprojects/model/user"
    "io/ioutil"
    "net/http"
    "os"
    "encoding/json"
)

func newUser(response http.ResponseWriter, request *http.Request) {
  defer request.Body.Close()
  bodySave, _ := ioutil.ReadAll(request.Body)
  userReceived := user.User{}
  json.Unmarshal(bodySave, &userReceived)
  id := userReceived.Save()
  jsonValue := map[string]string{"Result": "OK"}
  jsonValue["id"] = id
  result, _ := json.Marshal(jsonValue)
  fmt.Fprintf(response, string(result))
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