package app

import (
	"imagedownloader/internal/config"
	"io"
	"net/url"
	"strings"
)

type downloader interface {
	Download(*url.URL) (io.ReadCloser, error)
}

type saver interface {
	SaveImages(io.ReadCloser, string) error
}

type imageParser interface {
	GetImageLinks(body io.ReadCloser) ([]*url.URL, error)
}

func DownloadImages(
	dl downloader,
	sv saver,
	pr imageParser,
	cfg *config.Config,
) error {
	respBody, err := dl.Download(cfg.Url())
	if err != nil {
		return err
	}
	defer respBody.Close()

	links, err := pr.GetImageLinks(respBody)
	if err != nil {
		return err
	}

	for _, link := range links {
		if strings.HasPrefix(link.String(), "/") {
			link.Scheme = cfg.Url().Scheme
			link.Host = cfg.Url().Host
		}

		image, err := dl.Download(link)
		if err != nil {
			return err
		}
		defer image.Close()

		splittedPath := strings.Split(link.Path, "/")
		sv.SaveImages(image, splittedPath[len(splittedPath)-1])
	}

	return nil
}
