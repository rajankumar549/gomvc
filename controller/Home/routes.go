package Home

import u "github.com/rajankumar549/Trim/utils"

var Export = []u.Action{
	u.Action{
		URL:         "/hash",
		Method:      "GET",
		Handler:     homeHandler,
		Description: "Hello world Request Handler",
	},
	u.Action{
		URL:         "/ping/{id:[0-9]+}",
		Method:      "GET",
		Handler:     pingHandller,
		Description: "Ping Pong Request Handler",
	},
	u.Action{
		URL:         "/link/short",
		Method:      "POST",
		Handler:     getShortURL,
		Description: "URL Shortner ",
	},
}
