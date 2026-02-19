# WebsiteCrawler Go SDK

The Go client for the [Website Crawler API](https://www.websitecrawler.org/apidoc) allows you to invoke the crawler remotely, poll the status of the crawl and retrieve the LLM ready JSON format data which can be either used to train LLMs or analyzing a website.

## Installation
```bash
go get github.com/pc8544/website-crawler-go-sdk
```

## Usage

Get your API key from [WebsiteCrawler.org](https://www.websitecrawler.org) settings page. Update the `examples/demo.go` file with your API key, URL/website to crawl and the crawl limit and run the demo. You will see the live status of the crawl and the current URL being crawled. Once the crawl job is over, a JSON will be returned. You can extract data from this JSON as per your requirements.
