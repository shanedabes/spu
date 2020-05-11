package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCallback(t *testing.T) {
	req, err := http.NewRequest("GET", "/callback?state=test&code=test", nil)
	assert.Nil(t, err)

	rr := httptest.NewRecorder()

	ch := make(chan string)
	handler := NewCallbackHandler("test", ch)

	go http.HandlerFunc(handler).ServeHTTP(rr, req)

	code := <-ch

	assert.Equal(t, code, "test")

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d .\n Got %d instead", http.StatusOK, status)
	}
}

func TestNewCallbackFormMismatch(t *testing.T) {
	req, err := http.NewRequest("GET", "/callback?state=test&code=test", nil)
	assert.Nil(t, err)

	rr := httptest.NewRecorder()

	ch := make(chan string)
	handler := NewCallbackHandler("wrong", ch)

	go http.HandlerFunc(handler).ServeHTTP(rr, req)

	code := <-ch

	assert.Equal(t, code, "error")
}
