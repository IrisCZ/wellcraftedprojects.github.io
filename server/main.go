package main

import (
  "github.com/jjballano/wellcraftedprojects/http"
  "github.com/jjballano/wellcraftedprojects/database/mongo"
  "github.com/jjballano/wellcraftedprojects/model/user"
    "os"
)

func main() {

    mongo.Init("localhost","wellcrafted")

    user.Init(new (mongo.Mongo))

    port := os.Getenv("PORT")
    if len(port) < 1 {
        port = "1337"
    }
    http.StartApi(port)
}