package filebrowserfs

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

func (f *FilebrowserFS) newRequest(method string, url *url.URL, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(
		method,
		url.String(),
		body,
	)
	if err != nil {
		return nil, err
	}

	req.AddCookie(&http.Cookie{
		Name:  "auth",
		Value: f.token,
	})

	req.Header.Add("X-Auth", f.token)

	return req, nil
}

func responseJsonBody(res *http.Response, out any) error {
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, out)
	if err != nil {
		return err
	}

	return nil
}

func responseTextBody(res *http.Response) (string, error) {
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func doAndClose(req *http.Request) (*http.Response, error) {
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	res.Body.Close()

	return res, nil
}

func do(req *http.Request) (*http.Response, error) {
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func addQueryBool(url *url.URL, name string, value bool) {
	addQuery(
		url,
		name,
		strconv.FormatBool(value),
	)
}

func addQuery(url *url.URL, name string, value string) {
	query := url.Query()
	query.Add(name, value)
	url.RawQuery = query.Encode()
}
