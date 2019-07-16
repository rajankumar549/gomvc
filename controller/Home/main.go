package Home

type t struct {
	Name string `json:name`
}

func homeHandler(id string, params ...interface{}) (interface{}, error) {
	return t{
		Name: "Rajan Kumar",
	}, nil
}
func pingHandller(id string, params ...interface{}) (interface{}, error) {
	return t{
		Name: "Rajan Kumar",
	}, nil
}
