package Home

import u "github.com/rajankumar549/Trim/utils"

var Export = []u.Action{
	u.Action{
		URL:         "/",
		Method:      "GET",
		Handler:     homeHandler,
		Description: "Hello world Request Handler",
	},
	u.Action{
		URL:         "/ping",
		Method:      "GET",
		Handler:     pingHandller,
		Description: "Ping Pong Request Handler",
	},
}
