package controllers_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"demowebgo/controllers"
	"demowebgo/config"
)

func TestCreateUser(t *testing.T) {
	config.ConnectDatabase()
	router := setupRouter()

	body := `{
		"name": "Hieu",
		"email": "hieu@gmail.com",
		"password": "123456"
	}`

	req, _ := http.NewRequest("POST", "/user", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
	assert.Contains(t, w.Body.String(), `"email":"hieu@gmail.com"`)
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	r.POST("/user", controllers.CreateUser)

	return r
}