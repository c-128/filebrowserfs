package filebrowserfs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type loginRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Recaptcha string `json:"recaptcha"`
}

// Log into a File Browser server for the given username and password
// You must call client.RenewToken() before the token expires (default is 2 hours)
func Login(baseURL string, username string, password string) (*FilebrowserFS, error) {
	url, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(&loginRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		return nil, err
	}

	loginUrl := url.JoinPath("/api/login").String()
	res, err := http.Post(
		loginUrl,
		"application/json",
		bytes.NewReader(data),
	)
	if err != nil {
		return nil, err
	}

	token, err := responseTextBody(res)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to login: %d", res.StatusCode)
	}

	filebrowser := &FilebrowserFS{
		baseURL: url,
		token:   token,
	}
	return filebrowser, nil
}
