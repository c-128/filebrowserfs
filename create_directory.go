package filebrowserfs

import (
	"fmt"
	"net/http"
)

// Creates a directory at the path
// It creats all directories if it is nested
func (client *FilebrowserFS) CreateDirectory(path string, override bool) error {
	url := client.baseURL.JoinPath("/api/resources", path, "/")
	addQueryBool(url, "override", override)

	req, err := client.newRequest(http.MethodPost, url, nil)
	if err != nil {
		return err
	}

	res, err := doAndClose(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to create directory: %d", res.StatusCode)
	}

	return nil
}
