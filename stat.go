package filebrowserfs

import (
	"net/http"
)

func (f *FilebrowserFS) Stat(path string) (*FileInfo, error) {
	url := f.baseURL.JoinPath("/api/resources", path)

	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	var out *FileInfo
	err = responseJsonBody(res, &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}
