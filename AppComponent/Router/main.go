package APPRouter

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	_MUX "github.com/gorilla/mux"
	Console "github.com/rajankumar549/Trim/AppComponent/Console"
	_Controller "github.com/rajankumar549/Trim/controller"
	u "github.com/rajankumar549/Trim/utils"
)

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

func RouterInit() *_MUX.Router {
	r := _MUX.NewRouter()
	r.Use(loggingMiddleware)
	for k, _ := range _Controller.Export {
		Action := _Controller.Export[k]
		fmt.Printf("URL : %+v Description : %+v Type : %+v\n", Action.URL, Action.Description, Action.Method)
		r.HandleFunc(Action.URL, func(w http.ResponseWriter, r *http.Request) {
			ServeHTTP(w, r, Action.Handler)
			return
		})
	}
	return r
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Console.DEBUG("URL => %s", r.URL)
		next.ServeHTTP(w, r)
	})
}

func ServeHTTP(w http.ResponseWriter, r *http.Request, method u.ActionHandler) {
	response := Response{}
	response.Base.Status = "OK"

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET,PUT,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "X-Tkpd-UserId,Tkpd-UserId,Authorization,Origin")

	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}

	start := time.Now()

	var data interface{}
	var err error
	vars := _MUX.Vars(r)
	query := r.URL.Query()

	for k, v := range query {
		v = u.Unique(v)
		vars[k] = strings.Join(v, ",")
	}
	decoder := json.NewDecoder(r.Body)
	data, err = method(vars, decoder)

	response.Base.ServerProcessTime = time.Since(start).String()

	var buf []byte

	w.Header().Set("Content-Type", "application/json")

	if data != nil && err == nil {
		response.Data = data
		if buf, err = json.Marshal(response); err == nil {
			Console.INFO(string(buf[:]))
			w.Write(buf)
			return
		}
	}

	buf, _ = json.Marshal(response)
	Console.INFO(string(buf[:]))
	w.Write(buf)
	return
}
