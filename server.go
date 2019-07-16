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
	Router := _APPRouter.RouterInit()
	port := cfg.GetPORT()
	err := http.ListenAndServe(port, Router)
	if err != nil {
		Console.EMERGENCY("Server Failed %v , %s", err, "tstrdcgw")
		os.Exit(1)
	}
	Console.INFO("Server Statring At Port %s", port)
}
