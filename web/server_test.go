package web

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReturns200OnRoot(t *testing.T) {
	mockRender := func(w http.ResponseWriter, size int) {

	}
	mux := Server(mockRender)
	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	mux.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Error("Expect 200 but got ", recorder.Code)
	}
}

func TestSetsGridSizeAccordingToParam(t *testing.T) {
	var receivedSize int
	mockRender := func(w http.ResponseWriter, size int) {
		receivedSize = size
	}
	mux := Server(mockRender)
	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/?size=20", nil)

	mux.ServeHTTP(recorder, req)

	if receivedSize != 20 {
		t.Error("Expected size to be 20 but got ", receivedSize)
	}
}
