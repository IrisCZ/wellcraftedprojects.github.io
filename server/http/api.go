package http

import (
    "fmt"
    "github.com/gorilla/mux"
    "net/http"
    "encoding/json"
)

func notFound(response http.ResponseWriter, request *http.Request) {
  fmt.Fprintf(response, "Error")
}

func StartApi(port string) {
  router := mux.NewRouter()
  router.HandleFunc("/user/new", NewUser).Methods("POST")
  router.HandleFunc("/login", Login).Methods("POST")
  router.HandleFunc("/", notFound)
  http.Handle("/", router)
  http.ListenAndServe(":"+port, nil)
}

func getEntity(r *http.Request, v Entity) error {
    return v.UnmarshalHTTP(r)
}

func parseResponseTo(response http.ResponseWriter, result string, params map[string]interface{}){
    params["Result"] = result
    resultJson, _ := json.Marshal(params)

    fmt.Fprintf(response, string(resultJson))
}