package filebrowserfs

import (
	"fmt"
	"net/http"
)

func (f *FilebrowserFS) Rename(
	oldpath string,
	newpath string,
	override bool,
	rename bool,
) error {
	url := f.baseURL.JoinPath("/api/resources", oldpath)
	addQuery(url, "action", "rename")
	addQuery(url, "destination", newpath)
	addQueryBool(url, "override", override)
	addQueryBool(url, "rename", rename)

	req, err := f.newRequest(http.MethodPatch, url, nil)
	if err != nil {
		return err
	}

	res, err := doAndClose(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to rename path: %d", res.StatusCode)
	}

	return nil
}
