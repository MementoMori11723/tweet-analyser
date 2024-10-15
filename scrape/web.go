package scrape

import (
	"context"
	"log"

	"github.com/chromedp/chromedp"
)

func GetTweets(url string) []string {
  ctx, cancel := chromedp.NewContext(context.Background())
  defer cancel()
  if ctx.Err() != nil {
    log.Fatal(ctx.Err())
  }
  return []string{
    "tweet1",
    "tweet2",
  }
}

func GetTweet(url string) string {
  return ""
}
