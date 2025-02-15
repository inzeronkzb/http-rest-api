package apiserver

import (
	"testing"
  	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"net/http"
)
func TestAPIServer_HandleHello(t *testing.T){
	s := New(NewConfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	s.handleHello().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "Hello")
}