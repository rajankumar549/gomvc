package APPRouter

import (
	"fmt"
	"net/http"

	_MUX "github.com/gorilla/mux"
	Console "github.com/rajankumar549/Trim/AppComponent/Console"
	_Controller "github.com/rajankumar549/Trim/controller"
)

func RouterInit() *_MUX.Router {
	r := _MUX.NewRouter()
	r.Use(loggingMiddleware)
	for k, _ := range _Controller.Export {
		Action := _Controller.Export[k]
		fmt.Printf("URL : %+v Description : %+v Type : %+v\n", Action.URL, Action.Description, Action.Method)
		r.HandleFunc(Action.URL, Action.Handller)
	}
	return r
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Console.DEBUG("URL => %s", r.URL)
		next.ServeHTTP(w, r)
	})
}
