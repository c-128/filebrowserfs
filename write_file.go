package filebrowserfs

import (
	"fmt"
	"io"
	"net/http"
)

// Write to a file
// This overwrites the file and does not create a new one if it does not exist
func (f *FilebrowserFS) WriteFile(path string, reader io.Reader) error {
	url := f.baseURL.JoinPath("/api/resources", path)

	req, err := f.newRequest(http.MethodPut, url, reader)
	if err != nil {
		return err
	}

	res, err := doAndClose(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to open file: %d", res.StatusCode)
	}

	return nil
}
