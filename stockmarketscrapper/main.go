package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

// company name, price, change in price string

type Stock struct {
	company, price, change string
}

func main() {
	ticker := []string{
		"BHARTIARTL.NS", // Bharti Airtel Limited
		"RELIANCE.NS",   // Reliance Industries Limited
		"TECHM.NS",      // Tech Mahindra Limited
		"INDUSINDBK.NS", // IndusInd Bank Limited
		"BAJAJ-AUTO.NS", // Bajaj Auto Limited
		"CIPLA.NS",      // Cipla Limited
		"TCS.NS",        // Tata Consultancy Services Limited
		"NTPC.NS",       // NTPC Limited
		"TATASTEEL.NS",  // Tata Steel Limited
		"ULTRACEMCO.NS", // UltraTech Cement Limited
		"KOTAKBANK.NS",  // Kotak Mahindra Bank Limited
		"LT.NS",         // Larsen & Toubro Limited
		"MARUTI.NS",     // Maruti Suzuki India Limited
		"HDFCLIFE.NS",   // HDFC Life Insurance Company Limited
		"WIPRO.NS",      // Wipro Limited
		"TATACONSUM.NS", // Tata Consumer Products Limited
		"TITAN.NS",      // Titan Company Limited
		"ONGC.NS",       // Oil and Natural Gas Corporation Limited
		"HEROMOTOCO.NS", // Hero MotoCorp Limited
		"BAJFINANCE.NS", // Bajaj Finance Limited
	}
	stocks := []Stock{}

	// initialize the colly

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting:", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Error occured:", err)
	})

	c.OnHTML("div#quote-header-info", func(e *colly.HTMLElement) {
		stock := Stock{}
		stock.company = e.ChildText("h1")
		fmt.Println("Company:", stock.company)
		//stock.price = e.ChildText("fin-streamer[data-field='regularMarketPrice]")
		stock.price = e.ChildText("fin-streamer[data-field='regularMarketPrice']")
		fmt.Println("Price:", stock.price)
		//stock.change = e.ChildText("fin-streamer[data-field='regularMarketChange")
		stock.change = e.ChildText("fin-streamer[data-field='regularMarketChange']")
		fmt.Println("Change:", stock.change)

		stocks = append(stocks, stock)
	})
	c.Wait()

	for _, t := range ticker {
		c.Visit("https://finance.yahoo.com/quote/" + t + "/")
	}
	fmt.Println(stocks)

	file, err := os.Create("stocks.csv")
	if err != nil {
		log.Fatalln("failed to create", err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	headers := []string{
		"company",
		"price",
		"change",
	}
	writer.Write(headers)
	for _, stock := range stocks {
		record := []string{
			stock.company,
			stock.price,
			stock.change,
		}
		writer.Write(record)
		defer writer.Flush()
	}

}
