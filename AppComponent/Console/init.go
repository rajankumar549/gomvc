package Console

import (
	"fmt"
	"time"
)

type LogLevel struct {
	Level int64
	Desc  string
}

var Settings = map[int64]string{
	100: "DEBUG",
	200: "INFO",
	250: "NOTICE",
	300: "WARNING",
	400: "ERROR",
	500: "CRITICAL",
	550: "ALERT",
	600: "EMERGENCY",
}

var config LogLevel

func InitLogger(level int64) {

	if val, ok := Settings[level]; ok {
		config.Desc = val
		config.Level = level
		return
	}
	fmt.Printf("Invalid Log Level")

}
func addRecord(level int64, format string, a ...interface{}) {
	if level >= config.Level {
		t := time.Now()
		format = "\n[" + Settings[level] + "] : [" + t.Format("2006-01-02T15:04:05Z07:00") + "] : " + format
		fmt.Println(len(a))
		fmt.Println(a)
		if len(a) == 0 {
			fmt.Println(len(a))
			fmt.Print(format)
			return
		}
		fmt.Printf(format, a)
	}
}
func DEBUG(format string, a ...interface{}) {
	addRecord(100, format, a)
}
func INFO(format string, a ...interface{}) {
	addRecord(200, format, a)
}
func NOTICE(format string, a ...interface{}) {
	addRecord(250, format, a)
}
func WARNING(format string, a ...interface{}) {
	addRecord(300, format, a)
}
func ERROR(format string, a ...interface{}) {
	addRecord(400, format, a)
}
func CRITICAL(format string, a ...interface{}) {
	addRecord(500, format, a)
}
func ALERT(format string, a ...interface{}) {
	addRecord(550, format, a)
}
func EMERGENCY(format string, a ...interface{}) {
	addRecord(600, format, a)
}
