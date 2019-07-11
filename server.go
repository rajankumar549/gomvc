package main

import (
	"net/http"

	Console "github.com/rajankumar549/Trim/AppComponent/Console"
	_APPRouter "github.com/rajankumar549/Trim/AppComponent/Router"
	cfg "github.com/rajankumar549/Trim/config"
)

func main() {
	Console.InitLogger(100)
	Router := _APPRouter.RouterInit()
	http.ListenAndServe(cfg.GetPORT(), Router)
}
