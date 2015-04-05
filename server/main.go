package main

import (
  "github.com/IrisCZ/wellcraftedprojects/http"
  "github.com/IrisCZ/wellcraftedprojects/database/mongo"
  "github.com/IrisCZ/wellcraftedprojects/model/user"
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