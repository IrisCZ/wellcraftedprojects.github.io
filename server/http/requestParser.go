package http

import (
  "io"
  "io/ioutil"
  "encoding/json"
    "fmt"
    "net/http"
)

func ParseBodyTo(body io.ReadCloser, model interface{}){
  bodySave, _ := ioutil.ReadAll(body)
  json.Unmarshal(bodySave, &model)
}

func ParseResponseTo(response http.ResponseWriter, result string, params []ResponseParam){
  jsonValue := map[string]string{"Result": result}
  length := len(params)
  for i:=0; i < length; i++ {
    jsonValue[params[i].Name] = params[i].Value
  }
  resultJson, _ := json.Marshal(jsonValue)
  fmt.Fprintf(response, string(resultJson))
}
