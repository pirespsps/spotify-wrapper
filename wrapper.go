package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
)

func getClientID() (string, error) {

	var rurl = "https://accounts.spotify.com/api/token"

	data := url.Values{
		"grant_type":    {"client_credentials"},
		"client_id":     {ID},
		"client_secret": {Secret},
	}

	reader := bytes.NewBufferString(data.Encode())

	req, err := http.NewRequest("POST", rurl, reader)
	if err != nil {
		return "", errors.New("client token requisition error")
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", errors.New("client token response error")
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", errors.New("client token body error")
	}

	token := struct {
		Token string `json:"access_token"`
		Type  string `json:"token_type"`
	}{}

	json.Unmarshal(body, &token)

	return token.Token, nil

}

func genericRequest(url string) ([]byte, error) {
	token, err := getClientID()
	if err != nil {
		return nil, errors.New("generic request token error")
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.New("generic request new request error")
	}
	req.Header.Add("Authorization", "Bearer "+token)

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.New("generic request response error")
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errors.New("generic request body error")
	}

	return body, nil
}

func makeAction() {

}
