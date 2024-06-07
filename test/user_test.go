package tests

import (
	"net/http"
	"net/http/httptest"
	"referral-system-2/controllers"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	router := gin.Default()
	router.POST("/api/register", controllers.Register)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/register", strings.NewReader(`{"email": "test@example.com", "password": "password"}`))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
	assert.Contains(t, w.Body.String(), "test@example.com")
}

func TestLogin(t *testing.T) {
	router := gin.Default()
	router.POST("/api/login", controllers.Login)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/login", strings.NewReader(`{"email": "test@example.com", "password": "password"}`))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "token")
}
