package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func Test_GET_returns_StatusNotFound(t *testing.T) {
	t.Log("Expect GET to return status code 404")
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	if err != nil {
		t.Error(err)
	}

	resRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(HashHandler)
	handler.ServeHTTP(resRecorder, req)

	//Check status code is 404
	status := resRecorder.Code
	if status != http.StatusNotFound {
		t.Errorf("Expected %v but instead got %v", http.StatusNotFound, status)
	}

	//Check body is Bad Request
	expectedMessage := "Not Found"
	message := resRecorder.Body.String()
	if message != expectedMessage {
		t.Errorf("Expected %v but instead got %v", expectedMessage, message)
	}
}

func Test_POST_bad_form_returns_StatusBadRequest(t *testing.T) {
	t.Log("Expect bad form in POST to return status code 400")
	formData := url.Values{}
	formData.Set("pass", "foo")
	req, err := http.NewRequest(http.MethodPost, "http://localhost:8080", strings.NewReader(formData.Encode()))
	if err != nil {
		t.Error(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(HashHandler)
	handler.ServeHTTP(resRecorder, req)

	//Check status code is 404
	status := resRecorder.Code
	if status != http.StatusBadRequest {
		t.Errorf("Expected %v but instead got %v", http.StatusBadRequest, status)
	}

	//Check body is Bad Request
	expectedMessage := "Bad Request"
	message := resRecorder.Body.String()
	if message != expectedMessage {
		t.Errorf("Expected %v but instead got %v", expectedMessage, message)
	}
}

func Test_password_hash_is_correct(t *testing.T) {
	t.Log("Expect password hash to be correct")
	expectedHashedPass := "ZEHhWB65gUlzdVwtDQArEyx+KVLzp/aTaRaPlBzYRIFj6vjFdqEb0Q5B8zVKCZ0vKbZPZklJz0Fd7su2A+gf7Q=="
	formData := url.Values{}
	formData.Add("password", "angryMonkey")

	req, err := http.NewRequest(http.MethodPost, "http://localhost:8080", strings.NewReader(formData.Encode()))
	if err != nil {
		t.Error(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(HashHandler)
	handler.ServeHTTP(resRecorder, req)

	//Check status code is 200
	status := resRecorder.Code
	if status != http.StatusOK {
		t.Errorf("Expected %v but instead got %v", http.StatusOK, status)
	}

	//Check body is hash
	hashedPass := resRecorder.Body.String()
	if hashedPass != expectedHashedPass {
		t.Errorf("Expected %v but instead got %v", expectedHashedPass, hashedPass)
	}
}
