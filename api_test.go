package main

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleStuff(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/stuff?filename=test.txt", nil)
	HandleStuff(w, req)
	body, _ := ioutil.ReadAll(w.Result().Body)
	expected := `{
		"count": 2,
		"total": 2,
		"page": 0,
		"page_size": 100,
		"data": [
			{"id": 1, "name": "Fred"},
			{"id": 2, "name": "Bob"}
		]
	}
	`
	assert.JSONEq(t, expected, string(body))
}

func TestHandleStuffError(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/foobar?filename=doesnotexist", nil)
	HandleStuff(w, req)
	body, _ := ioutil.ReadAll(w.Result().Body)
	expected := `{
		"error": "open doesnotexist: no such file or directory"
	}
	`
	assert.JSONEq(t, expected, string(body))
}
