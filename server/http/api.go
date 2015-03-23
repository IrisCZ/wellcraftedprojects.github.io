package http

import (
    "fmt"
    "github.com/gorilla/mux"
    "github.com/jjballano/wellcraftedprojects/model/user"
    "net/http"
    "os"
    "encoding/json"
)

type ResponseParam struct {
    Name string
    Value string
}

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

func getEntity(r *http.Request, v Entity) error {
    return v.UnmarshalHTTP(r)
}

func parseResponseTo(response http.ResponseWriter, result string, params []ResponseParam){
    jsonValue := map[string]string{"Result": result}
    length := len(params)
    for i:=0; i < length; i++ {
        jsonValue[params[i].Name] = params[i].Value
    }
    resultJson, _ := json.Marshal(jsonValue)
    fmt.Fprintf(response, string(resultJson))
}