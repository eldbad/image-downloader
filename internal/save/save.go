package save

import (
	"io"
	"os"
)

type FileSystemSaver struct{}

func (s FileSystemSaver) SaveImages(body io.ReadCloser, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, body)
	if err != nil {
		return err
	}

	return nil
}
