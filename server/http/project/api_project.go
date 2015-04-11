package project

import (
  "net/http"
  "github.com/IrisCZ/wellcraftedprojects/model/project"
  "github.com/IrisCZ/wellcraftedprojects/http/utils"
  "encoding/json"
    "fmt"
)

func List(response http.ResponseWriter, request *http.Request){

  var list []interface{}
  projects := project.FindAll()
  if projects != nil {
    for _,element := range projects {
      result,_ := json.Marshal(element)
      aProject := new(project.Project)
      json.Unmarshal(result,aProject)
      list = append(list, *aProject)
    }
  }
  utils.ParseResponseTo(response,"",map[string]interface{}{"projects":list})
}

func New(response http.ResponseWriter, request *http.Request){
  fmt.Println("LLEGAAA")

    theProject,_ := parseProject(request)
    aa,_ := json.Marshal(theProject)
    fmt.Println("VALOR", string(aa ))
}


func parseProject(request *http.Request)(*project.Project, error){
  theProject := new(project.Project)
  return theProject, utils.GetEntity(request, theProject)
}
