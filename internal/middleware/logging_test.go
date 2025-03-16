package middleware

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLogginMiddleware(t *testing.T) {
	var buf bytes.Buffer

	log.SetOutput(&buf)

	mockHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("test response"))
	})

	loggingHandler := LoggingMiddleware(mockHandler)

	req, err := http.NewRequest("GET", "/test-path", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.RemoteAddr = "127.0.0.1:1234"

	rr := httptest.NewRecorder()
	loggingHandler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler retornou código de status incorreto: got %v want %v", status, http.StatusOK)
	}

	if rr.Body.String() != "test response" {
		t.Errorf("Handler retornou corpo incorreto: got %v want %v", rr.Body.String(), "test response")
	}

	logOutput := buf.String()

	if !strings.Contains(logOutput, "GET") {
		t.Errorf("Log nao contém o método HTTP: %s", logOutput)
	}

	if !strings.Contains(logOutput, "/test-path") {
		t.Errorf("Log nao contém o caminho da requisição: %s", logOutput)
	}

	if !strings.Contains(logOutput, "127.0.0.1:1234") {
		t.Errorf("Log nao contém o endereço remoto: %s", logOutput)
	}
}
