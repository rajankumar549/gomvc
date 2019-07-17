package Utils

import (
	json "encoding/json"
	"net/http"
)

type HandlerFunc func(rw http.ResponseWriter, r *http.Request) (interface{}, error)
type ActionHandler func(vars map[string]string, params *json.Decoder) (interface{}, error)
type Action struct {
	URL         string
	Handler     ActionHandler
	Method      string `default:"GET"`
	Description string
}

func Unique(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
