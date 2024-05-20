package transdevclient

type SearchResult struct {
	Stops  []Stop  `json:"stops"`
	Routes []Route `json:"routes"`
}

type SearchResponse struct {
	StatusCode int `json:"statusCode"`
	Result     struct {
		Results SearchResult `json:"results"`
	} `json:"result"`
}
