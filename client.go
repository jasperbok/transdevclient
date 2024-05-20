package transdevclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	client *http.Client
}

func NewClient() Client {
	httpClient := http.DefaultClient
	return Client{client: httpClient}
}

// Search searches for Stops and Routes with query in their name. It
// returns a SearchResult containing matching Stops and Routes.
//
// Note that only a limited amount of objects is returned in the
// SearchResult. Make your query as specific as possible for a higher
// probability to find the objects you need.
func (c *Client) Search(query string) (SearchResult, error) {
	req, err := http.NewRequest(http.MethodGet, "https://www.breng.nl/api/travelplanner/search", nil)
	if err != nil {
		return SearchResult{}, err
	}

	q := req.URL.Query()
	q.Add("text", query)
	req.URL.RawQuery = q.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return SearchResult{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return SearchResult{}, err
	}

	searchResponse := SearchResponse{}
	err = json.Unmarshal(body, &searchResponse)
	if err != nil {
		return SearchResult{}, err
	}

	return searchResponse.Result.Results, nil
}

func (c *Client) Departures(quayID, name, town string) ([]Departure, error) {
	req, err := http.NewRequest(http.MethodGet, "https://www.breng.nl/api/travelplanner/quays/departures", nil)
	if err != nil {
		return []Departure{}, err
	}

	q := req.URL.Query()
	q.Add("stop", fmt.Sprintf("{'quayid': '%s', 'name': '%s', 'town': '%s'}", quayID, name, town))
	req.URL.RawQuery = q.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return []Departure{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []Departure{}, err
	}

	departuresResponse := DeparturesResponse{}
	err = json.Unmarshal(body, &departuresResponse)
	if err != nil {
		return []Departure{}, err
	}

	if len(departuresResponse.Results) == 0 {
		return []Departure{}, errors.New("no departures found")
	}

	if len(departuresResponse.Results) > 1 {
		return []Departure{}, errors.New("too many departures found")
	}

	return departuresResponse.Results[0].Departures, nil
}
