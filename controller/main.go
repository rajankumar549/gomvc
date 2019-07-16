package Controller

import (
	_HomeController "github.com/rajankumar549/Trim/Controller/Home"
	u "github.com/rajankumar549/Trim/utils"
)

var Export = make(map[int]u.Action)

func init() {
	var actions = make(map[int]u.Action)
	for k, v := range _HomeController.Export {
		actions[k] = v
	}
	Export = actions
}
