package main

import (
    "context"
    "time"
    "log"
    "github.com/chromedp/chromedp"
)

func main() {

  opts := append(chromedp.DefaultExecAllocatorOptions[:],
    chromedp.Flag("headless", true),
    chromedp.Flag("remote-debugging-port", "9222"),
  )

  allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
  defer cancel()

  ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
  defer cancel()

  chromedp.Run(ctx,
  chromedp.Sleep(time.Minute),)


}
