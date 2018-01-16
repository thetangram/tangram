package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthyHandler(t *testing.T) {
	handlerSUT := healthyHandler
	r, _ := http.NewRequest("GET", "/healthy", nil)
	if err := handlerTester(r, handlerSUT, healthyStatus); err != nil {
		t.Errorf("healthy handler error %v", err)
	}
	if err := handlerTester(r, handlerSUT, healthyPayload); err != nil {
		t.Errorf("healthy handler error %v", err)
	}
}

/*
 WARNING! To be able to test this feature, the application status probably needs to
        be injected or mocked
*/
func TestReadyHandler(t *testing.T) {
	handlerSUT := readyHandler
	r, _ := http.NewRequest("GET", "/ready", nil)
	if err := handlerTester(r, handlerSUT, notReadyStatus); err != nil {
		t.Errorf("ready handler error %v", err)
	}
	if err := handlerTester(r, handlerSUT, notReadyPayload); err != nil {
		t.Errorf("ready handler error %v", err)
	}
}

func BenchmarkReadyHandler(b *testing.B) {
	r, _ := http.NewRequest("GET", "/ready", nil)
	for i := 0; i < b.N; i++ {
		handlerTester(r, readyHandler, func(rr *httptest.ResponseRecorder) (err error) { return nil })
	}
}

func BenchmarkHealthyHandler(b *testing.B) {
	r, _ := http.NewRequest("GET", "/ready", nil)
	for i := 0; i < b.N; i++ {
		handlerTester(r, healthyHandler, func(rr *httptest.ResponseRecorder) (err error) { return nil })

	}
}

func healthyStatus(rr *httptest.ResponseRecorder) (err error) {
	if status := rr.Code; status != http.StatusOK {
		return fmt.Errorf("healthy handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	return
}

func healthyPayload(rr *httptest.ResponseRecorder) (err error) {
	expected := "OK"
	if rr.Body.String() != expected {
		return fmt.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
	return
}

func notReadyStatus(rr *httptest.ResponseRecorder) (err error) {
	if status := rr.Code; status != http.StatusServiceUnavailable {
		return fmt.Errorf("healthy handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	return
}

func notReadyPayload(rr *httptest.ResponseRecorder) (err error) {
	expected := "NO"
	if rr.Body.String() != expected {
		return fmt.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
	return
}

func handlerTester(r *http.Request,
	h func(w http.ResponseWriter, r *http.Request),
	v func(rr *httptest.ResponseRecorder) error) error {
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h)
	handler.ServeHTTP(rr, r)
	return v(rr)
}
