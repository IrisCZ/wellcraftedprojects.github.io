package project

import (
  "net/http"
  "github.com/IrisCZ/wellcraftedprojects/model/project"
  "github.com/IrisCZ/wellcraftedprojects/http/utils"
  "encoding/json"
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
  theProject,error := parseProject(request)
  if error != nil {
    returnError(response, error)
    return
  }

  id,error := theProject.Save()
  if error != nil {
    returnError(response, error)
  } else {
    utils.ParseResponseTo(response, "OK", map[string]interface{}{"id":id})
  }
}


func parseProject(request *http.Request)(theProject *project.Project, error error){
  theProject = new(project.Project)
  error = utils.GetEntity(request, theProject)
  return
}

func returnError(response http.ResponseWriter, error error){
    utils.ParseResponseTo(response, "", map[string]interface{}{"error":error.Error()})
}
