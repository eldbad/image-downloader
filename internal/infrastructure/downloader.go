package infrastructure

import (
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Downloader struct{}

func (d *Downloader) DownloadDocument(link *url.URL) (io.ReadCloser, error) {
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
