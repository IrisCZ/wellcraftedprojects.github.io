package http

import (
  "fmt"
  "github.com/gorilla/mux"
  "net/http"
  "github.com/IrisCZ/wellcraftedprojects/http/user"
  "github.com/IrisCZ/wellcraftedprojects/http/project"
)

func notFound(response http.ResponseWriter, request *http.Request) {
  fmt.Fprintf(response, "Error")
}

func StartApi(port string) {
  router := mux.NewRouter()
  router.HandleFunc("/user/new", user.New).Methods("POST")
  router.HandleFunc("/login", user.Login).Methods("POST")
  router.HandleFunc("/projects", project.List).Methods("GET")
  router.HandleFunc("/project", project.New).Methods("POST")
  router.HandleFunc("/", notFound)
  http.Handle("/", &CORS{router})
  http.ListenAndServe(":"+port, nil)
}

type CORS struct {
  router *mux.Router
}


func (s *CORS) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
    if origin := req.Header.Get("Origin"); origin != "" {
        rw.Header().Set("Access-Control-Allow-Origin", origin)
        rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
        rw.Header().Set("Access-Control-Allow-Headers",
        "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
    }
    // Stop here if its Preflighted OPTIONS request
    if req.Method == "OPTIONS" {
        return
    }
    s.router.ServeHTTP(rw, req)
}