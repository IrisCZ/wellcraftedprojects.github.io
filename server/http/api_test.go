package http


import (
  "testing"
  "github.com/jjballano/wellcraftedprojects/database/mongo"
  "github.com/jjballano/wellcraftedprojects/model/user"
  "github.com/jjballano/wellcraftedprojects/model"
  "github.com/jjballano/wellcraftedprojects/crypto"
  "os"
  "net/http"
  "net/http/httptest"
  "bytes"
  "github.com/stretchr/testify/assert"
)

type MongoMock struct{
}

var mongoMock MongoMock
var responseExpected string

func (mongo MongoMock) Save(obj model.Model, collectionName string) string {
  theUser := obj.(*user.User)
  responseExpected = `{"email":"`+theUser.Email+`","password":"`+theUser.Password+`"}`
  return responseExpected
}

func TestMain(m *testing.M){
  mongo.Init("localhost","wellcrafted")
  mongoMock = MongoMock{}
  user.Init(mongoMock)

  retCode := m.Run()

  os.Exit(retCode)

}

func Test_asks_mongo_for_saving_a_user_with_params_received(t *testing.T) {

  handler := func(response http.ResponseWriter, request *http.Request) {
    newUser(response,request)
  }
  reader := bytes.NewReader([]byte(`{"email": "an@email.com", "password":"anyPassword"}`))
  req, _ := http.NewRequest("POST", "/user/new", reader)

  recoder := httptest.NewRecorder()
  handler(recoder, req)

  passwordExpected := crypto.EncodePassword("anyPassword")

  assert.Equal(t, responseExpected, `{"email":"an@email.com","password":"`+passwordExpected+`"}`, "They should be equal")

}

