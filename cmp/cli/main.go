package main

import (
	"imagedownloader/internal/app"
	"imagedownloader/internal/config"
	"imagedownloader/internal/download"
	"imagedownloader/internal/parse"
	"imagedownloader/internal/save"
	"log"
	"os"
)

func main() {
	cfg, err := config.NewConfig(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	dl := download.HttpDownloader{}
	sv := save.FileSystemSaver{}
	pr := parse.HtmlParser{}

	err = app.DownloadImages(&dl, &sv, &pr, cfg)
	if err != nil {
		log.Fatal(err)
	}
}
