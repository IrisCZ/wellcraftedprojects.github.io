package main

import (
  "github.com/jjballano/wellcraftedprojects/http"
  "github.com/jjballano/wellcraftedprojects/database/mongo"
  "github.com/jjballano/wellcraftedprojects/model/user"
)

func main() {

    mongo.Init("localhost","wellcrafted")

    user.Init(mongo.Mongo{})

    http.StartApi()
}