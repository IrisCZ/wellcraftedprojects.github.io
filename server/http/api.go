package http

import (
    "fmt"
    "github.com/gorilla/mux"
    "net/http"
    "os"
    "encoding/json"
)

type ResponseParam struct {
    Name string
    Value string
}


func notFound(response http.ResponseWriter, request *http.Request) {
  fmt.Fprintf(response, "Error")
}

func StartApi() {
  router := mux.NewRouter()
  router.HandleFunc("/user/new", NewUser).Methods("POST")
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