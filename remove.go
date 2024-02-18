package filebrowserfs

import (
	"fmt"
	"net/http"
)

// Deletes a file/directory
// The directory does not need to be empty
func (f *FilebrowserFS) Remove(path string) error {
	url := f.baseURL.JoinPath("/api/resources", path)

	req, err := f.newRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	res, err := doAndClose(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to remove path: %d", res.StatusCode)
	}

	return nil
}
