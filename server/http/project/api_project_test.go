package project

import(
  "github.com/IrisCZ/wellcraftedprojects/database/mongo"
  "github.com/IrisCZ/wellcraftedprojects/model"
  "github.com/IrisCZ/wellcraftedprojects/model/project"
  "testing"
  "net/http"
  "os"
  "net/http/httptest"
  "encoding/json"
  "github.com/stretchr/testify/assert"
)

type MongoMock struct{
}


func (mongo MongoMock) Save(obj model.Model, collectionName string) (string, error) {
    return "ID",nil
}

func (mongo MongoMock) FindOne(collectionName string, params map[string]string, model model.Model) error{
    return nil
}

func (mongo MongoMock) FindAll(collectionName string) ([]map[string]interface{}, error) {
    result := make([]map[string]interface{},2)
    result[0] = map[string]interface{}{"url":"anyUrl", "author":"any author", "description":"any description", "positives":5, "negatives":2, "image":"/any/image.jpg"}
    result[1] = map[string]interface{}{"url":"anotherUrl"}
    return result,nil
}

func TestMain(m *testing.M){
    mongo.Init("localhost","wellcrafted")
    project.Init(MongoMock{})

    retCode := m.Run()

    os.Exit(retCode)
}

var handler = func(response http.ResponseWriter, request *http.Request) {
    if(request.URL.String() == "/projects"){
        List(response,request)
    }
}

func Test_ask_mongo_for_project_list(t *testing.T){
  req, _ := http.NewRequest("GET", "/projects", nil)

  recoder := httptest.NewRecorder()
  handler(recoder, req)

  response := parseResponse(recoder)

  assert.Equal(t, response["error"], nil)

  projects := []project.Project{}
  projectsReturned := response["projects"].([]interface{})
  for _,element := range projectsReturned {
    result,_ := json.Marshal(element)
    aProject := new(project.Project)
    json.Unmarshal(result,aProject)
    projects = append(projects, *aProject)
  }

  assert.Equal(t, len(projects), 2)
  assert.Equal(t, projects[0].Author, "any author")
  assert.Equal(t, projects[0].Url, "anyUrl")
  assert.Equal(t, projects[0].Description, "any description")
  assert.Equal(t, projects[0].Positives, 5)
  assert.Equal(t, projects[0].Negatives, 2)
  assert.Equal(t, projects[0].Image, "/any/image.jpg")
  assert.Equal(t, projects[1].Url, "anotherUrl")
}

func parseResponse(recoder *httptest.ResponseRecorder) map[string]interface{} {
    var response map[string]interface{}
    json.Unmarshal(recoder.Body.Bytes(), &response)
    return response
}