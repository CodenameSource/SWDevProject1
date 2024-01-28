package scraper

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

func extractNumericChars(input string) string {
	re := regexp.MustCompile("[0-9.]+")
	numericChars := re.FindString(input)

	return numericChars
}

func CheckAmazonPrice(url string) (float64, error) {
	var price float64
	var priceStr string

	collector := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64; rv:10.0) Gecko/20100101 Firefox/10.0"),
	)

	collector.OnHTML("span.aok-align-center:nth-child(1) > span:nth-child(1)", func(e *colly.HTMLElement) {
		priceStr = e.Text
	})

	collector.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	err := collector.Visit(url)
	if err != nil {
		return 0, err
	}

	// Clean up the extracted price (remove extra spaces and newlines)
	priceStr = strings.TrimSpace(priceStr)

	fmt.Println("Scraped price ", priceStr)

	price, err = strconv.ParseFloat(extractNumericChars(priceStr), 64)

	if err != nil {
		return 0, err
	}

	return price, nil
}
