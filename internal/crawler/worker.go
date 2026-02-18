package crawler

import (
	"Homework2/internal/images"
	"Homework2/internal/storage"
	"Homework2/internal/web"
	"context"
	"fmt"
	"path/filepath"
)

func (e *Engine) worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			return
		case url, ok := <-e.jobs:
			if !ok {
				return
			}

			// Skip if already visited (safe due to URLSet's internal mutex)
			if !e.visited.Add(url) {
				continue
			}

			// Acquire semaphore slot
			e.sem <- struct{}{}
			// Launch the actual job in a new goroutine
			go func(u string) {
				// Ensure the slot is released when the job finishes.
				defer func() { <-e.sem }()

				normalized, err := NormalizeURL(u)
				if err != nil {
					return
				}

				if !ShouldVisit(u, normalized, e.FollowExternal) {
					return
				}

				// Chromedp fetch page
				html, err := FetchPage(ctx, normalized)
				if err != nil {
					return
				}

				// Extract and process images
				imgs := images.ExtractImages(normalized, html)
				for _, img := range imgs {
					path, err := images.DownloadImage(img, "assets/originals")
					if err != nil {
						continue
					}

					fmt.Printf("[worker %d] Downloaded image %s\n", id, path)

					var thumbPath string
					if img.Format == "svg" {
						thumbPath = path
					} else {
						thumbPath, err = images.CreateThumbnail(path, "assets/thumbnails")
						if err != nil {
							continue
						}
					}

					var meta *images.Metadata
					if img.Format == "svg" {
						meta = &images.Metadata{
							Width:  0,
							Height: 0,
						}
					} else {
						meta, err = images.ExtractMetadata(path)
						if err != nil {
							continue
						}
					}

					record := storage.ImageRecord{
						URL:           img.URL,
						Filename:      filepath.Base(path),
						Alt:           img.Alt,
						Title:         img.Title,
						Width:         meta.Width,
						Height:        meta.Height,
						Format:        img.Format,
						ThumbnailPath: web.ToWebPath(thumbPath),
					}

					// Safe because e.DB is initialized in main
					e.DB.Insert(record)
				}

				// Extract links and enqueue safely
				links := ExtractLinks(normalized, html)
				for _, link := range links {
					normLink, err := NormalizeURL(link)
					if err != nil {
						continue
					}

					if !ShouldVisit(normalized, normLink, e.FollowExternal) {
						continue
					}

					// Context-safe send to jobs channel
					select {
					case <-ctx.Done():
						return
					case e.jobs <- normLink:
					}
				}
			}(url)
		}
	}
}
