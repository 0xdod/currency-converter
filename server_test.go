package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGETConvert(t *testing.T) {
	t.Run("should perform successful conversion", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/convert/10/ksh/ghs", nil)
		response := httptest.NewRecorder()

		conversionResult := ConversionResult{
			From: Currency{"ksh", 10},
			To:   Currency{"ghs", 0.55},
		}

		server := NewServer()
		server.ServeHTTP(response, request)

		want := buildResponse("success", "conversion successful", &conversionResult)
		got := &APIResponse{}

		_ = json.NewDecoder(response.Body).Decode(&got)

		if response.Code != http.StatusOK {
			t.Errorf("Expeected %d status code, but got %d", http.StatusOK, response.Code)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %+v, want %+v", got, want)
		}
	})

	t.Run("should return error", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/convert/fyi", nil)
		response := httptest.NewRecorder()

		server := &Server{}

		server.ServeHTTP(response, request)

		got := &APIResponse{}

		_ = json.NewDecoder(response.Body).Decode(got)

		if response.Code != http.StatusBadRequest {
			t.Errorf("Expeected %d status code, but got %d", http.StatusBadRequest, response.Code)
		}

		if got.Status != "error" {
			t.Errorf("got %s status, expected %s status", got.Status, "error")
		}
	})
}
