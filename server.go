package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Server struct {
}

// APIResponse represents the data send back to the user by the api.
type APIResponse struct {
	Status  string           `json:"status,omitempty"`
	Message string           `json:"message,omitempty"`
	Data    ConversionResult `json:"data"`
}

func buildResponse(status, message string, data ConversionResult) *APIResponse {
	return &APIResponse{status, message, data}
}

// JSON converts the custom APIResponse type to valid JSON
func (a *APIResponse) JSON() string {
	data, err := json.Marshal(a)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

// /convert/amount/from/to
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	paths := strings.Split(strings.TrimPrefix(r.URL.Path, "/convert/"), "/")

	if len(paths) < 3 {
		w.WriteHeader(401)
		response := buildResponse("Error", "Bad request", ConversionResult{})
		fmt.Fprint(w, response.JSON())
		return
	}

	amount, _ := strconv.ParseFloat(paths[0], 64)
	from := paths[1]
	to := paths[2]

	res := Convert(from, to, amount)

	response := buildResponse("", "", res)

	fmt.Fprint(w, response.JSON())
}
