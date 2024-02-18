package filebrowserfs

import (
	"net/http"
)

func (f *FilebrowserFS) ReadDirectory(path string) (*DirectoryListing, error) {
	url := f.baseURL.JoinPath("/api/resources", path)

	req, err := f.newRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	var out *DirectoryListing
	err = responseJsonBody(res, &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}
