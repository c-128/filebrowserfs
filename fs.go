package filebrowserfs

import (
	"net/url"
)

type FilebrowserFS struct {
	baseURL *url.URL
	token   string
}
