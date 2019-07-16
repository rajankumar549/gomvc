package Console

import (
	"fmt"
	"time"

	"github.com/fatih/color"
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
func addRecord(c color.Attribute, level int64, format string, a ...interface{}) {
	if level >= config.Level {
		t := time.Now()
		timeStapText := "[" + Settings[level] + "] : [" + t.Format("2006-01-02T15:04:05Z07:00") + "] : "
		logText := format
		switch level {
		case 600:
			color.Set(c, color.Underline, color.Bold)
			break
		case 550:
			color.Set(c, color.Bold)
			break
		case 500:
			color.Set(c, color.Italic)
			break
		default:
			color.Set(c)
		}

		if len(a) > 0 {
			logText = fmt.Sprintf(format, a...)
		}
		fmt.Println(timeStapText, logText)
		color.Unset()
	}
}
func DEBUG(format string, a ...interface{}) {
	addRecord(color.FgHiWhite, 100, format, a...)
}
func INFO(format string, a ...interface{}) {
	addRecord(color.FgGreen, 200, format, a...)
}
func NOTICE(format string, a ...interface{}) {
	addRecord(color.FgCyan, 250, format, a...)
}
func WARNING(format string, a ...interface{}) {
	addRecord(color.FgYellow, 300, format, a...)
}
func ERROR(format string, a ...interface{}) {
	addRecord(color.FgMagenta, 400, format, a...)
}
func CRITICAL(format string, a ...interface{}) {
	addRecord(color.FgRed, 500, format, a...)
}
func ALERT(format string, a ...interface{}) {
	addRecord(color.FgHiRed, 550, format, a...)
}
func EMERGENCY(format string, a ...interface{}) {
	addRecord(color.FgWhite, 600, format, a...)
}
