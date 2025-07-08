package web_test

import (
	"ascii-art-web-stylize/server"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMainHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.MainHandler)
	handler.ServeHTTP(rr, req)

	// Check the status code is 200
	if rr.Code != http.StatusOK {
		t.Errorf("MainHandler returned wrong status code: got %v, want %v", rr.Code, http.StatusOK)
	}

	// Check the response contains expected content
	if !bytes.Contains(rr.Body.Bytes(), []byte("<title>ASCII Art Generator</title>")) {
		t.Errorf("MainHandler response does not contain expected content")
	}
}

func TestMainHandlerInvalidPath(t *testing.T) {
	req, err := http.NewRequest("GET", "/invalid-path", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.MainHandler)
	handler.ServeHTTP(rr, req)

	// Check the status code is 404
	if rr.Code != http.StatusNotFound {
		t.Errorf("MainHandler returned wrong status code: got %v, want %v", rr.Code, http.StatusNotFound)
	}
}

func TestResultHandlerValidInput(t *testing.T) {
	form := "input-text=Hello&banner=standard"
	req, err := http.NewRequest("POST", "/error", bytes.NewBufferString(form))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.ResultHandler)
	handler.ServeHTTP(rr, req)

	// Check the status code is 200
	if rr.Code != http.StatusOK {
		t.Errorf("ResultHandler returned wrong status code: got %v, want %v", rr.Code, http.StatusOK)
	}

	// Check if response contains ASCII Art
	if rr.Body.Len() == 0 {
		t.Errorf("ResultHandler returned an empty body")
	}
}

func TestResultHandlerInvalidInput(t *testing.T) {
	form := "input-text=Invalid\x01Text&banner=standard"
	req, err := http.NewRequest("POST", "/error", bytes.NewBufferString(form))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.ResultHandler)
	handler.ServeHTTP(rr, req)

	// Check the status code is 400
	if rr.Code != http.StatusBadRequest {
		t.Errorf("ResultHandler returned wrong status code: got %v, want %v", rr.Code, http.StatusBadRequest)
	}
}

func TestResultHandlerInvalidBanner(t *testing.T) {
	form := "input-text=Hello&banner=invalid-banner"
	req, err := http.NewRequest("POST", "/error", bytes.NewBufferString(form))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.ResultHandler)
	handler.ServeHTTP(rr, req)

	// Check the status code is 404
	if rr.Code != http.StatusNotFound {
		t.Errorf("ResultHandler returned wrong status code: got %v, want %v", rr.Code, http.StatusNotFound)
	}
}
