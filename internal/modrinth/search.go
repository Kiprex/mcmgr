package modrinth

import (
	"net/url"
	"net/http"
	"encoding/json"
	"io"
)


func Search(query string) (SearchResponse, error) {
	v := url.Values{}
	v.Set("query", query)
	resp, err := http.Get("https://api.modrinth.com/v2/search?" + v.Encode())
	if err != nil {
		return SearchResponse{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return SearchResponse{}, err
	}
	var result SearchResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return SearchResponse{}, err
	}
	return result, nil
}
