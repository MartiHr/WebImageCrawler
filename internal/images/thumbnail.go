package images

import (
	"path/filepath"

	"github.com/disintegration/imaging"
)

func CreateThumbnail(srcPath, dstDir string) (string, error) {
	img, err := imaging.Open(srcPath)
	if err != nil {
		return "", err
	}

	thumb := imaging.Resize(img, 200, 0, imaging.Lanczos)

	filename := filepath.Base(srcPath)
	dstPath := filepath.Join(dstDir, filename)

	err = imaging.Save(thumb, dstPath)
	return dstPath, err
}
