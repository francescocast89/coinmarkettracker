package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// CryptoMarketMap structure
type CryptoMarketMap struct {
	CryptoMarket map[string]*CryptoMarket `json:"data"`
}

// CryptoMarket structure
type CryptoMarket struct {
	ID                int               `json:"id"`
	Name              string            `json:"name"`
	Symbol            string            `json:"symbol"`
	Slug              string            `json:"slug"`
	CirculatingSupply float64           `json:"circulating_supply"`
	TotalSupply       float64           `json:"total_supply"`
	MaxSupply         float64           `json:"max_supply"`
	DateAdded         string            `json:"date_added"`
	NumMarketPairs    int               `json:"num_market_pairs"`
	CMCRank           int               `json:"cmc_rank"`
	LastUpdated       string            `json:"last_updated"`
	Quote             map[string]*Quote `json:"quote"`
}

type Quote struct {
	Price            float64 `json:"price,omitempty"`
	Volume24H        float64 `json:"volume_24h,omitempty"`
	Volume7D         float64 `json:"volume_7d,omitempty"`
	Volume30D        float64 `json:"volume_30d,omitempty"`
	Volume24Hbase    float64 `json:"volume_24h_base,omitempty"`
	Volume24Hquote   float64 `json:"volume_24h_quote,omitempty"`
	PercentChange1H  float64 `json:"percent_change_1h,omitempty"`
	PercentChange24H float64 `json:"percent_change_24h,omitempty"`
	PercentChange7D  float64 `json:"percent_change_7d,omitempty"`
	PercentChange30D float64 `json:"percent_change_30d,omitempty"`
	MarketCap        float64 `json:"market_cap,omitempty"`
	LastUpdated      string  `json:"last_updated"`
}

func main() {
	observedTokens := []string{"BTC", "ETC", "ADA", "DOT", "MATIC"}
	//observedTokens := []string{"BTC"}

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := url.Values{}
	q.Add("symbol", strings.Join(observedTokens[:], ","))
	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", "64295071-9fd2-4522-8e39-194d20bb341f")
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request to server")
		os.Exit(1)
	}
	fmt.Println(resp.Status)
	respBody, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(respBody))
	var response CryptoMarketMap
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		fmt.Println(err)
	}
	for _, i := range response.CryptoMarket {
		fmt.Println(i.Name, " ", i.Quote["USD"].Price)
	}
}
