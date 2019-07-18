package Home

type GetShortURLRequest struct {
	URL string `json:"url"`
}

type GetShortURLResponse struct {
	URL       string `json:"url"`
	HostName  string `json:"host"`
	Id        string `json:"host"`
	ValidUpto time   `json:vaild_upto`
}
