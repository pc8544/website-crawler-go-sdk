package main

import (
    "fmt"
    "log"
    "time"
    "github.com/pc8544/website-crawler-go-sdk"
)

func main() {
    url := "YOUR_URL"   //replace YOUR_URL by a url/domain starting with https:// or http://
    limit := 10     //replace 10 with the number of URLs you want WebsiteCrawler.org to crawl

    client := websitecrawler.NewClient("YOUR_API_KEY") //replace YOUR_API_KEY with the WebsiteCrawler API key
     if err := client.Authenticate(); err != nil {
        log.Fatal("Auth failed:", err)
    }

    _, err := client.Start(url, limit)
    if err != nil {
        log.Fatal("Start failed:", err)
    }

    status, err := client.PollStatus(url, 5*time.Second)
    if err != nil {
        log.Fatal("Polling failed:", err)
    }
    fmt.Println("Final Status:", status)

    data, err := client.Data(url)
    if err != nil {
        log.Fatal("Data fetch failed:", err)
    }
    fmt.Println("Crawl Data:", data.Pages)
}
