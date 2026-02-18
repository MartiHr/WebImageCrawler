package crawler

import (
	"context"
	"time"

	"github.com/chromedp/chromedp"
)

func FetchPage(ctx context.Context, url string) (string, error) {
	// Create isolated browser context
	ctx, cancel := chromedp.NewContext(ctx)
	defer cancel()

	// Safety timeout per page
	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	var html string

	err := chromedp.Run(ctx,
		chromedp.Navigate(url),

		// Wait until DOM is ready
		chromedp.WaitReady("body", chromedp.ByQuery),

		// Give JS time to render (important for React/Vue)
		chromedp.Sleep(2*time.Second),

		// Extract final DOM
		chromedp.OuterHTML("html", &html),
	)

	return html, err
}
