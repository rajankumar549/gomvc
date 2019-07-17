package Home

import (
	"encoding/json"

	u "github.com/rajankumar549/Trim/utils"
)

type t struct {
	Name string      `json:name`
	Data interface{} `json:data`
}

func homeHandler(vars map[string]string, params *json.Decoder) (interface{}, error) {
	url := vars["url"]
	v, err := u.GetHash([]byte(url))
	if err != nil {
		//console.ERROR("Error : %+v", err)
		return t{
			Name: "Error",
			Data: err,
		}, nil

	}
	return t{
		Name: "Has of " + url,
		Data: v,
	}, nil

}
func pingHandller(vars map[string]string, params *json.Decoder) (interface{}, error) {
	return t{
		Name: "Rajan Kumar",
		Data: vars,
	}, nil
}
