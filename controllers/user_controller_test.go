package controllers

import (
	"encoding/xml"
	"github/publiusvergilius/mongodb-course/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Controller (c *gin.Context) {
	c.XML(http.StatusCreated, "")	
} 

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/users", Controller)

	return r
}

/** 
*  !TODO: Usar o UserController
*  !TODO: Criar StubMongoDB
*/
func TestCreateUserXML(t *testing.T){
	router := setupRouter()
	user := models.User{ID: primitive.NewObjectID(), Username: "Vini", Email: "test@email.com"}

	out, _ := xml.MarshalIndent(user, "", " ")

	req := httptest.NewRequest("POST", "/users", strings.NewReader(string(out)))
	req.Header.Set("Content-Type", "application/xml")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "<Name>Vini</Name>")
	assert.Contains(t, w.Body.String(), "<Email>test@email.com</Email>")
}
