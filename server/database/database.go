package database

import "github.com/jjballano/wellcraftedprojects/model"

type Database interface{
  Save(obj model.Model, collectionName string) string
}
