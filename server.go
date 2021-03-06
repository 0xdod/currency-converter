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
	http.Handler
}

func NewServer() *Server {
	s := &Server{}
	router := http.NewServeMux()
	router.HandleFunc("/convert/", convertHandler)
	s.Handler = router
	return s
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
		log.Println(err)
		return ""
	}
	return string(data)
}

func convertHandler(w http.ResponseWriter, r *http.Request) {
	paths := strings.Split(strings.TrimPrefix(r.URL.Path, "/convert/"), "/")
	if len(paths) < 3 {
		w.WriteHeader(http.StatusBadRequest)
		response := buildResponse("error", "Bad request", nil)
		fmt.Fprint(w, response.JSON())
		return
	}
	amount, _ := strconv.ParseFloat(paths[0], 64)
	from := paths[1]
	to := paths[2]

	res, err := Convert(from, to, amount)

	if err != nil {
		w.WriteHeader(422)
		response := buildResponse("error", err.Error(), &res)
		fmt.Fprint(w, response.JSON())
		return
	}

	response := buildResponse("success", "conversion succcessful", &res)

	fmt.Fprint(w, response.JSON())
}
