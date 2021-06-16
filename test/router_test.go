package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestPingRouter(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	rep, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, rep)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
