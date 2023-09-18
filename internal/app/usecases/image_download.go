package usecases

import (
	"imagedownloader/internal/app/parse"
	"imagedownloader/internal/domain"
	"io"
	"net/url"
	"strings"
)

type downloader interface {
	DownloadDocument(*url.URL) (io.ReadCloser, error)
}

type saver interface {
	SaveFile(io.ReadCloser, string) error
}

type ImageDownload struct{}

func (i ImageDownload) DownloadImages(
	downloader downloader,
	saver saver,
	config domain.Config,
) error {
	respBody, err := downloader.DownloadDocument(config.Url())
	if err != nil {
		return err
	}
	defer respBody.Close()

	parser := parse.Parser{}
	links, err := parser.GetImageLinks(respBody)
	if err != nil {
		return err
	}

	for _, link := range links {
		if strings.HasPrefix(link.String(), "/") {
			link.Scheme = config.Url().Scheme
			link.Host = config.Url().Host
		}

		image, err := downloader.DownloadDocument(link)
		if err != nil {
			return err
		}
		defer image.Close()

		splittedPath := strings.Split(link.Path, "/")
		saver.SaveFile(image, splittedPath[len(splittedPath)-1])
	}

	return nil
}
