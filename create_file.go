package filebrowserfs

import (
	"fmt"
	"net/http"
)

// Creats a file at the path
// Does not write to it
func (f *FilebrowserFS) CreateFile(path string, override bool) error {
	url := f.baseURL.JoinPath("/api/resources", path)
	addQueryBool(url, "override", override)

	req, err := f.newRequest(http.MethodPost, url, nil)
	if err != nil {
		return err
	}

	res, err := doAndClose(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to create file: %d", res.StatusCode)
	}

	return nil
}
