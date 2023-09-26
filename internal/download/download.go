package download

import (
	"io"
	"net/http"
	"net/url"
	"strings"
)

type HttpDownloader struct{}

func (d HttpDownloader) Download(link *url.URL) (io.ReadCloser, error) {
	strL := link.String()
	if strings.HasPrefix(strL, "/") {
		strL = ""
	}
	resp, err := http.Get(strL)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}
