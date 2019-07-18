package Home

import (
	"encoding/json"
	"errors"
	"time"

	u "github.com/rajankumar549/Trim/utils"
	TrimRedis "github.com/rajankumar549/Trim/utils/redis"
)

type t struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
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

func getShortURL(vars map[string]string, params *json.Decoder) (interface{}, error) {
	data := GetShortURLRequest{}
	if err := params.Decode(&data); err != nil || len(data.URL) == 0 {
		return nil, errors.New("Invalid Request")
	}

	hash, err := u.GetHash([]byte(data.URL))
	if err != nil {
		return nil, errors.New("Internal Server Error")
	}

	result, err := TrimRedis.RedisClient.SetWithNX(string(hash[:]), data.URL, 86400)
	if result != "1" {
		return nil, errors.New("URL Already Exist")
	}
	return GetShortURLResponse{
		URL:       "http://localhost:8012/" + hash,
		HostName:  "localhost",
		Id:        hash,
		ValidUpto: time.Now(),
	}, nil

}
