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

var paramsReceived = make([] interface{}, 2)

func (mongo MongoMock) Save(obj model.Model, collectionName string) (string, error) {
  paramsReceived[0] = obj.(*user.User)
  paramsReceived[1] = collectionName
  return "ID",nil
}

var handler = func(response http.ResponseWriter, request *http.Request) {
    newUser(response,request)
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

    assert.Equal(t, recoder.Body.String(), `{"Result":"OK","id":"ID"}`)
}



