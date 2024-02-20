package filebrowserfs

import "net/url"

// Create a new client using a token
// You must call client.RenewToken() before the token expires (default is 2 hours)
func NewClient(baseURL string, token string) (*FilebrowserFS, error) {
	url, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	filebrowser := &FilebrowserFS{
		baseURL: url,
		token:   token,
	}
	return filebrowser, nil
}
