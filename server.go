package main

import (
	"net/http"
	"os"

	Console "github.com/rajankumar549/Trim/AppComponent/Console"
	_APPRouter "github.com/rajankumar549/Trim/AppComponent/Router"
	cfg "github.com/rajankumar549/Trim/config"
)

func main() {

	Console.InitLogger(100)
	Console.NOTICE("App Starting....")
	Router := _APPRouter.RouterInit()
	port := cfg.GetPORT()
	Console.INFO("Server at Port %s", port)
	err := http.ListenAndServe(port, Router)
	if err != nil {
		Console.EMERGENCY("Server Failed %v , %s", err)
		os.Exit(1)
	}

}
