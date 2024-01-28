package kafkamessenger

import (
	"fmt"
	"log"

	"vot-hw1-checkprices/scraper"
)

func ExtractPrice(url string) error {
	price, err := scraper.CheckAmazonPrice(url)

	if err != nil {
		log.Println("Error checking Amazon price:", err)
		return err
	}

	fmt.Println("Scraped price ", price, " from ", url)

	err = SendRefreshedPrice(url, price)

	if err != nil {
		log.Println("Error sending refreshed price:", err)
		return err
	}

	return nil
}
