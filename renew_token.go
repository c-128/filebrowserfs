package filebrowserfs

import (
	"fmt"
	"net/http"
)

// Renew the token
// By default a token expires after 2 hours
func (f *FilebrowserFS) RenewToken() error {
	url := f.baseURL.JoinPath("/api/renew")
	req, err := f.newRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to renew token: %d", res.StatusCode)
	}

	newToken, err := responseTextBody(res)
	if err != nil {
		return err
	}

	f.token = newToken
	return nil
}
