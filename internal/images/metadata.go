package images

import (
	"os"

	"image"
	_ "image/jpeg"
	_ "image/png"
	//_ "image/webp"
	// Third-party decoder for WebP
	_ "golang.org/x/image/webp"
)

type Metadata struct {
	Width  int
	Height int
}

func ExtractMetadata(path string) (*Metadata, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	cfg, _, err := image.DecodeConfig(f)
	if err != nil {
		return nil, err
	}

	return &Metadata{
		Width:  cfg.Width,
		Height: cfg.Height,
	}, nil
}
