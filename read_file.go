package filebrowserfs

import (
	"fmt"
	"io"
	"net/http"
)

func (f *FilebrowserFS) ReadFile(path string) (io.ReadCloser, error) {
	url := f.baseURL.JoinPath("/api/raw", path)

	req, err := f.newRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to open file: %d", res.StatusCode)
	}

	return res.Body, nil
}
