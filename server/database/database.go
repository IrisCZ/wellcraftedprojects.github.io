package database

import "github.com/IrisCZ/wellcraftedprojects/model"

type Database interface{
  Save(obj model.Model, collectionName string) (string, error)
  FindOne(collectionName string, params map[string]string, obj model.Model) error
}
