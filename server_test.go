package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETConvert(t *testing.T) {
	t.Run("returns conversion result", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/convert/10/ksh/ghs", nil)
		response := httptest.NewRecorder()

		server := &Server{}

		server.ServeHTTP(response, request)

		got := 0.55

		apiResponse := &APIResponse{}

		data, _ := io.ReadAll(response.Body)

		_ = json.Unmarshal(data, apiResponse)

		want := apiResponse.Data["to"].Value

		if got != want {
			t.Errorf("got %.2f, want %.2f", got, want)
		}
	})
}
