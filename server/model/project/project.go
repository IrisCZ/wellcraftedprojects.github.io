package project

import(
  "github.com/IrisCZ/wellcraftedprojects/database"
  "gopkg.in/mgo.v2/bson"
  "encoding/json"
  "net/http"
  "io/ioutil"
  "strings"
)

const collectionName string = "projects"

type Project struct {
  Id bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
  Url string `json:"url,omitempty" bson:"url,omitempty"`
  Name string `json:"name,omitempty" bson:"name,omitempty"`
  Author string `json:"author,omitempty" bson:"author,omitempty"`
  Description string `json:"description,omitempty" bson:"description,omitempty"`
  Image string `json:"image,omitempty" bson:"image,omitempty"`
  Tags []string `json:"tags,omitempty" bson:"tags,omitempty"`
  Positives int `json:"positives,omitempty" bson:"positives,omitempty"`
  Negatives int  `json:"negatives,omitempty" bson:"negatives,omitempty"`
}

type RawProject struct {
  *Project
  RawTags string `json:"tags"`
}


var db database.Database

func Init(database database.Database){
  db = database
}

func (project Project) SetId(id bson.ObjectId){
  project.Id = id
}


func (project *Project) Save() (string,error){
    newProject,_ := db.Save(project, collectionName)
    return newProject,nil
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
    rawProject := new(RawProject)
    error := json.Unmarshal(bodySave, rawProject)
    if error != nil{
        return error
    }
    *project = *rawProject.Project

    if len(rawProject.RawTags) > 0 {
      project.Tags = strings.Split(rawProject.RawTags, ",")
    }
    if !strings.HasPrefix(project.Url,"http://") && !strings.HasPrefix(project.Url,"https://") {
      project.Url = "http://"+project.Url
    }

    return nil
}