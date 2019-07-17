package Config

import "time"

var config = map[string]string{
	"APPNAME":    "Trim App",
	"PORT":       "8012",
	"APIPREFIX":  "",
	"DEBUGLEVEL": "1",
}
var T = config

type Redis map[string]*RedisConfig
type RedisConfig struct {
	Endpoint  string
	Timeout   time.Duration
	MaxIdle   int
	MaxActive int
}

func GetPORT() string {
	return ":" + config["PORT"]
}

func GetRedisConfig() RedisConfig {
	return RedisConfig{
		Endpoint:  "localhost:6379",
		Timeout:   100,
		MaxIdle:   100,
		MaxActive: 100,
	}
}
