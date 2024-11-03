package scraper

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

// New creates a new scraping context and extracts data from the given URL
func New(url string) string {
	var data, by string

	dataXpath := "/html/body/div[1]/div/div/div[2]/main/div/div/div/div[1]/div/section/div/div/div[1]/div/div/article/div/div/div[3]/div[1]"
	byXpath := "/html/body/div[1]/div/div/div[2]/main/div/div/div/div[1]/div/section/div/div/div[1]/div/div/article/div/div/div[2]/div[2]/div/div/div[1]/div/div/div[1]/div/a/div/div[1]/span/span"

	// Create a chromedp context with headless flags
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),              // Run in headless mode
		chromedp.Flag("no-sandbox", true),            // Needed for running as root in Docker
		chromedp.Flag("disable-gpu", true),           // Disable GPU acceleration
		chromedp.Flag("disable-dev-shm-usage", true), // Use /tmp instead of /dev/shm
	)
	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// Create the chromedp context
	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	log.Println("Scraping:", url)
	log.Println("Creating context...")
	log.Println("Performing scraping...")

	err := chromedp.Run(
		ctx,
		chromedp.Navigate(url),
		chromedp.Sleep(2*time.Second),
		chromedp.Text(dataXpath, &data, chromedp.BySearch),
		chromedp.Text(byXpath, &by, chromedp.BySearch),
	)
	if err != nil {
		log.Fatalf("Scraping failed: %v", err)
	}

	log.Println("Scraping successful.")
	return fmt.Sprintf("%s by %s", data, by)
}
