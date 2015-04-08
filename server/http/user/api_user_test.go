package user


import (
  "testing"
  "github.com/IrisCZ/wellcraftedprojects/database/mongo"
  "github.com/IrisCZ/wellcraftedprojects/model/user"
  "github.com/IrisCZ/wellcraftedprojects/model"
  "github.com/IrisCZ/wellcraftedprojects/crypto"
  "os"
  "net/http"
  "net/http/httptest"
  "bytes"
  "github.com/stretchr/testify/assert"
  "errors"
  "crypto/sha256"
  "encoding/base64"
  "encoding/json"
)

type MongoMock struct{
}

var paramsReceived = make([] interface{}, 2)

func (mongo MongoMock) Save(obj model.Model, collectionName string) (string, error) {
  paramsReceived[0] = obj.(*user.User)
  paramsReceived[1] = collectionName
  return "ID",nil
}

func (mongo MongoMock) FindOne(collectionName string, params map[string]string, model model.Model) error{
  hash := sha256.New()
  hash.Write([]byte("anyPassword"))
  expectedPassword := base64.StdEncoding.EncodeToString(hash.Sum(nil))
  if params["email"] != "user@exists.com" {
    return errors.New("Not found")
  }
  json.Unmarshal([]byte(`{"email":"`+params["email"]+`","password":"`+expectedPassword+`"}`), model)
  return nil
}

func (mongo MongoMock) FindAll(collectionName string) ([]map[string]interface{}, error) {
  return nil,nil
}

var handler = func(response http.ResponseWriter, request *http.Request) {
  if(request.URL.String() == "/user/new"){
    New(response,request)
  }
  if(request.URL.String() == "/user/login"){
    Login(response,request)
  }
}


func TestMain(m *testing.M){
  mongo.Init("localhost","wellcrafted")
  user.Init(MongoMock{})

  retCode := m.Run()

  os.Exit(retCode)

}

func Test_asks_mongo_for_saving_a_user_with_params_received(t *testing.T) {

  reader := bytes.NewReader([]byte(`{"email": "an@email.com", "password":"anyPassword"}`))
  req, _ := http.NewRequest("POST", "/user/new", reader)

  handler(httptest.NewRecorder(), req)

  userExpected := user.User{Email:"an@email.com", Password:crypto.EncodePassword("anyPassword")}

  assert.Equal(t, paramsReceived[0], &userExpected)
  assert.Equal(t, paramsReceived[1], "users")
}

func Test_returns_the_id_of_the_new_user(t *testing.T) {

  reader := bytes.NewReader([]byte(`{"email": "an@email.com", "password":"anyPassword"}`))
  req, _ := http.NewRequest("POST", "/user/new", reader)

  recoder := httptest.NewRecorder()
  handler(recoder, req)

  response := parseResponse(recoder)

  assert.Equal(t, response["Result"], "OK")
  assert.Equal(t, response["id"], "ID")
}

func Test_returns_an_error_if_email_is_not_provider_when_ask_for_new_user(t *testing.T) {

  reader := bytes.NewReader([]byte(`{"password":"anyPassword"}`))
  req, _ := http.NewRequest("POST", "/user/new", reader)

  recoder := httptest.NewRecorder()
  handler(recoder, req)

  response := parseResponse(recoder)

  assert.Equal(t, response["Result"], "ERROR")
  assert.Equal(t, response["error"], "Email and password are mandatory")
}

func Test_returns_an_error_if_password_is_not_provider_when_ask_for_new_user(t *testing.T) {

  reader := bytes.NewReader([]byte(`{"email": "an@email.com"}`))
  req, _ := http.NewRequest("POST", "/user/new", reader)

  recoder := httptest.NewRecorder()
  handler(recoder, req)

  response := parseResponse(recoder)

  assert.Equal(t, response["Result"], "ERROR")
  assert.Equal(t, response["error"], "Email and password are mandatory")
}

func Test_returns_an_error_if_user_already_exists_when_ask_for_new_user(t *testing.T) {

    reader := bytes.NewReader([]byte(`{"email": "user@exists.com", "password":"anyPassword"}`))
    req, _ := http.NewRequest("POST", "/user/new", reader)

    recoder := httptest.NewRecorder()
    handler(recoder, req)

    response := parseResponse(recoder)

    assert.Equal(t, response["Result"], "ERROR")
    assert.Equal(t, response["error"], "user.exists.error")
}

func Test_returns_an_error_if_login_is_not_provider_when_ask_for_login(t *testing.T) {

  reader := bytes.NewReader([]byte(`{"password":"anyPassword"}`))
  req, _ := http.NewRequest("POST", "/user/login", reader)

  recoder := httptest.NewRecorder()
  handler(recoder, req)

  response := parseResponse(recoder)

  assert.Equal(t, response["Result"], "ERROR")
  assert.Equal(t, response["error"], "Login and password are mandatory")
}

func Test_returns_an_error_if_password_is_not_provider_when_ask_for_login(t *testing.T) {

  reader := bytes.NewReader([]byte(`{"login": "an@email.com"}`))
  req, _ := http.NewRequest("POST", "/user/login", reader)

  recoder := httptest.NewRecorder()
  handler(recoder, req)

  response := parseResponse(recoder)

  assert.Equal(t, response["Result"], "ERROR")
  assert.Equal(t, response["error"], "Login and password are mandatory")
}

func Test_returns_OK_when_login_is_correct(t *testing.T) {
  reader := bytes.NewReader([]byte(`{"login": "user@exists.com", "password":"anyPassword"}`))
  req, _ := http.NewRequest("POST", "/user/login", reader)

  recoder := httptest.NewRecorder()
  handler(recoder, req)

  response := parseResponse(recoder)

  assert.Equal(t, response["Result"], "OK")
}

func Test_returns_the_session_token_with_user_info_when_login_is_correct(t *testing.T) {
  reader := bytes.NewReader([]byte(`{"login": "user@exists.com", "password":"anyPassword"}`))
  req, _ := http.NewRequest("POST", "/user/login", reader)

  recoder := httptest.NewRecorder()
  handler(recoder, req)

  response := parseResponse(recoder)
  tokenExpected := crypto.CreateToken(map[string]interface{}{"login":"user@exists.com"})
  assert.Equal(t, response["token"], tokenExpected)
}

func parseResponse(recoder *httptest.ResponseRecorder) map[string]interface{} {
  var response map[string]interface{}
  json.Unmarshal(recoder.Body.Bytes(), &response)
  return response
}

