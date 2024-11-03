package scraper

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func New(url string) string {
	var data string
	var by string

	dataXpath := "/html/body/div[1]/div/div/div[2]/main/div/div/div/div[1]/div/section/div/div/div[1]/div/div/article/div/div/div[3]/div[1]"
	byXpath := "/html/body/div[1]/div/div/div[2]/main/div/div/div/div[1]/div/section/div/div/div[1]/div/div/article/div/div/div[2]/div[2]/div/div/div[1]/div/div/div[1]/div/a/div/div[1]/span/span"

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	log.Println("Scraping: ", url)
	log.Println("creating context")
	log.Println("Performing scraping")

	err := chromedp.Run(
		ctx,
		chromedp.Navigate(url),
		chromedp.Sleep(time.Second*2),
		chromedp.Text(dataXpath, &data, chromedp.BySearch),
		chromedp.Text(byXpath, &by, chromedp.BySearch),
	)

	log.Println("Scraping done")

	if err != nil {
		panic(err)
	}
	log.Println("scraping successfull")
	return data + " by " + by
}
