package main

import (
	"net/http"
	"os"

	"github.com/tokopedia/tkp-trigger/common/redis"

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

	redisConfig := cfg.GetRedisConfig()
	success := redis.New(redisConfig.Endpoint, redis.NetworkTCP, redis.Options{
		MaxActive: redisConfig.MaxActive,
		MaxIdle:   redisConfig.MaxIdle,
		Timeout:   redisConfig.Timeout,
		Wait:      true,
	})
	if success != nil {
		Console.INFO("Redis Connected.")
	} else {
		Console.EMERGENCY("Redis Connection failed!!!!")
	}
	err := http.ListenAndServe(port, Router)
	if err != nil {
		Console.EMERGENCY("Server Failed %v , %s", err)
		os.Exit(1)
	}

}
