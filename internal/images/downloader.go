package images

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func DownloadImage(img Image, dir string) (string, error) {
	resp, err := http.Get(img.URL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	filename := filepath.Base(img.URL)
	path := filepath.Join(dir, filename)

	out, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return path, err
}
