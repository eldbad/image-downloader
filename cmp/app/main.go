package main

import (
	"imagedownloader/internal/app/usecases"
	"imagedownloader/internal/domain"
	"imagedownloader/internal/infrastructure"
	"log"
	"os"
)

func main() {
	config, err := domain.NewConfig(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	downloader := infrastructure.Downloader{}
	saver := infrastructure.Saver{}

	imgDl := usecases.ImageDownload{}
	err = imgDl.DownloadImages(&downloader, &saver, *config)
	if err != nil {
		log.Fatal(err)
	}
}
