package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETConvert(t *testing.T) {
	t.Run("should perform successful conversion", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/convert/10/ksh/ghs", nil)
		response := httptest.NewRecorder()

		server := &Server{}

		server.ServeHTTP(response, request)

		apiResponse := &APIResponse{}

		data, _ := io.ReadAll(response.Body)

		_ = json.Unmarshal(data, apiResponse)

		want := 0.55
		got := apiResponse.Data.To.Value

		if response.Code != http.StatusOK {
			t.Errorf("Expeected %d status code, but got %d", http.StatusOK, response.Code)
		}

		if got != want {
			t.Errorf("got %.2f, want %.2f", got, want)
		}
	})

	t.Run("should return error", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/convert/fyi", nil)
		response := httptest.NewRecorder()

		server := &Server{}

		server.ServeHTTP(response, request)

		apiResponse := &APIResponse{}

		data, _ := io.ReadAll(response.Body)

		_ = json.Unmarshal(data, apiResponse)

		if response.Code != http.StatusBadRequest {
			t.Errorf("Expeected %d status code, but got %d", http.StatusBadRequest, response.Code)
		}

		if apiResponse.Status != "error" {
			t.Errorf("got %s status, expected %s status", apiResponse.Status, "error")
		}
	})
}
