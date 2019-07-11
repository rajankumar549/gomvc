package Utils

import "net/http"

type HanddlerFunc func(w http.ResponseWriter, r *http.Request)

type Action struct {
	URL         string
	Handller    HanddlerFunc
	Method      string `default:"GET"`
	Description string
}
