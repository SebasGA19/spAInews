package web

import (
	"bytes"
	"encoding/json"
	"github.com/SebasGA19/spAInews/api/internal/common"
	"github.com/SebasGA19/spAInews/api/internal/controller"
	"github.com/SebasGA19/spAInews/api/internal/tests"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"net/http"
	"net/http/httptest"
)

const (
	RedisHost = "127.0.0.1"
	RedisPort = "6379"
)

type TestEngine struct {
	Controller *controller.Controller
	Engine     *gin.Engine
}

func (te *TestEngine) Close() error {
	return te.Controller.Close()
}

func NewTestEngine() *TestEngine {
	sqlDB := common.ConnectSQL(
		"spainews",
		"spainews",
		"127.0.0.1",
		"3306",
		"spainews",
		"parseTime=true&multiStatements=true",
		true,
	)
	mongoCollection := common.ConnectMongo(
		"spainews",
		"spainews",
		"127.0.0.1",
		"27017",
		"spainews").Collection("news")
	sessions := common.ConnectRedis(RedisHost, RedisPort, 0)
	registrations := common.ConnectRedis(RedisHost, RedisPort, 1)
	confirmEmails := common.ConnectRedis(RedisHost, RedisPort, 2)
	resetPasswords := common.ConnectRedis(RedisHost, RedisPort, 3)

	smtp := common.ConnectSMTP(
		"dashboard@dev-spainews.co",
		"dashboard@dev-spainews.co",
		"password",
		"localhost",
		"25",
	)
	smtp.Dev = true
	c := controller.NewController(sqlDB, mongoCollection, sessions, registrations, confirmEmails, resetPasswords, smtp)
	tests.ClearDB(sqlDB)
	return &TestEngine{
		Controller: c,
		Engine:     NewEngine(c),
	}
}

func (te *TestEngine) doRequest(method, uri string, header map[string]string, obj any) *http.Response {
	if header == nil {
		header = map[string]string{}
	}
	var body io.Reader
	if obj != nil {
		j, marshalError := json.Marshal(obj)
		if marshalError != nil {
			panic(marshalError)
		}
		body = bytes.NewReader(j)
		header["Content-Type"] = "application/json"
	}
	req, err := http.NewRequest(method, uri, body)
	if err != nil {
		panic(err)
	}
	for key, value := range header {
		req.Header.Set(key, value)
	}
	w := httptest.NewRecorder()
	te.Engine.ServeHTTP(w, req)
	return w.Result()
}

func (te *TestEngine) Get(uri string, header map[string]string, obj any) *http.Response {
	return te.doRequest(http.MethodGet, uri, header, obj)
}

func (te *TestEngine) Post(uri string, header map[string]string, obj any) *http.Response {
	return te.doRequest(http.MethodPost, uri, header, obj)
}

func (te *TestEngine) Delete(uri string, header map[string]string, obj any) *http.Response {
	return te.doRequest(http.MethodDelete, uri, header, obj)
}

func (te *TestEngine) Put(uri string, header map[string]string, obj any) *http.Response {
	return te.doRequest(http.MethodPut, uri, header, obj)
}
