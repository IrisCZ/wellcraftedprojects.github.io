package utils
import (
  "net/http"
  "encoding/json"
  "fmt"
)


type Entity interface {
    UnmarshalHTTP(*http.Request) error
}

func GetEntity(r *http.Request, v Entity) error {
    return v.UnmarshalHTTP(r)
}

func ParseResponseTo(response http.ResponseWriter, result string, params map[string]interface{}){
    if len(result) > 0 {
      params["Result"] = result
    }
    resultJson, _ := json.Marshal(params)
    fmt.Fprintf(response, string(resultJson))
}
