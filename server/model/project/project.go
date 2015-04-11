package project

import(
  "github.com/IrisCZ/wellcraftedprojects/database"
  "gopkg.in/mgo.v2/bson"
  "encoding/json"
    "net/http"
    "io/ioutil"
)

const collectionName string = "projects"

type Project struct {
  Id bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
  Url string `json:"url,omitempty"`
  Name string `json:"name,omitempty"`
  Author string `json:"author,omitempty"`
  Description string `json:"description,omitempty"`
  Image string `json:"image,omitempty"`
  Tags string `json:"tags,omitempty"`
  Positives int `json:"positives,omitempty"`
  Negatives int  `json:"negatives,omitempty"`
}


var db database.Database

func Init(database database.Database){
  db = database
}

func (project Project) SetId(id bson.ObjectId){
  project.Id = id
}

func FindAll() []Project {

  result, error := db.FindAll(collectionName)
  if error != nil {
    return nil
  }
  projects := []Project{}
  for _,element := range result {
    project := new(Project)
    value,_ := json.Marshal(element)
    json.Unmarshal(value,project)
    projects = append(projects,*project)
  }

  return projects
}


func (project *Project) UnmarshalHTTP(request *http.Request) error {
    defer request.Body.Close()
    bodySave, _ := ioutil.ReadAll(request.Body)
    error := json.Unmarshal(bodySave, project)
    if error != nil{
        return error
    }
    return nil
}