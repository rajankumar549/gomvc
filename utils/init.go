package Utils

import "net/http"

type HandlerFunc func(rw http.ResponseWriter, r *http.Request) (interface{}, error)
type ActionHandler func(id string, params ...interface{}) (interface{}, error)
type Action struct {
	URL         string
	Handler     ActionHandler
	Method      string `default:"GET"`
	Description string
}
type Base struct {
	Status            string      `json:"status"`
	Config            interface{} `json:"config"`
	ServerProcessTime string      `json:"server_process_time"`
	ErrorMessage      []string    `json:"message_error,omitempty"`
	StatusMessage     []string    `json:"message_status,omitempty"`
}

type Response struct {
	Base
	Data interface{} `json:"data"`
}
