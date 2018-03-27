package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type UrlExpandResponse struct {
	RedirectTo string `json:"redirect_to"`
	PreviewUrl string `json:"preview_url"`
}

type UrlShrinkRequest struct {
	Url string `json:"url"`
}

type UrlShrinkResponse struct {
	ShortUrl string `json:"short_url"`
}

// ExpandUrl sends a request to the Big Earl API
// to retrieve the information about a given short url
func ExpandUrl(shortUrl string) (UrlExpandResponse, error) {
	var responseMessage UrlExpandResponse

	// Prepare the HTTP request
	req, err := http.NewRequest("GET", shortUrl, nil)
	if err != nil {
		return responseMessage, err
	}
	req.Header.Set("Accept", "application/json")

	// Send the request to the Big Earl API
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return responseMessage, err
	}
	defer resp.Body.Close()

	// Process the JSON response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return responseMessage, err
	}

	err = json.Unmarshal(body, &responseMessage)
	if err != nil {
		return responseMessage, err
	}

	return responseMessage, nil
}

// ShrinkUrl sends a request to the Big Earl API to create
// a short URL for the given url to shrink
func ShrinkUrl(apiUrl string, urlToShrink string) (UrlShrinkResponse, error) {
	var responseMessage UrlShrinkResponse

	// Prepare JSON payload to the API
	requestMessage := UrlShrinkRequest{
		Url: urlToShrink,
	}
	jsonStr, err := json.Marshal(requestMessage)
	if err != nil {
		return responseMessage, err
	}

	// Prepare the request to the API
	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// Send the request to the Big Earl API
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return responseMessage, err
	}
	defer resp.Body.Close()

	// Process the JSON response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return responseMessage, err
	}

	err = json.Unmarshal(body, &responseMessage)
	if err != nil {
		return responseMessage, err
	}

	return responseMessage, nil
}
