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
	Status  string            `json:"status,omitempty"`
	Message string            `json:"message,omitempty"`
	Data    *ConversionResult `json:"data"`
}

func buildResponse(status, message string, data *ConversionResult) *APIResponse {
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
	if !strings.HasPrefix(r.URL.Path, "/convert/") {
		w.WriteHeader(404)
		fmt.Fprint(w, `{"status": "error", "message": "route not found"}`)
		return
	}

	paths := strings.Split(strings.TrimPrefix(r.URL.Path, "/convert/"), "/")
	if len(paths) < 3 {
		w.WriteHeader(401)
		response := buildResponse("error", "Bad request", nil)
		fmt.Fprint(w, response.JSON())
		return
	}

	amount, _ := strconv.ParseFloat(paths[0], 64)
	from := paths[1]
	to := paths[2]

	res, err := Convert(from, to, amount)

	if err != nil {
		w.WriteHeader(403)
		response := buildResponse("error", err.Error(), &res)
		fmt.Fprint(w, response.JSON())
		return
	}

	response := buildResponse("success", "Conversion succcessful", &res)

	fmt.Fprint(w, response.JSON())
}
