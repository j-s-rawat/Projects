package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type PriceResponse struct {
	Price float64 `json:"price"`
}

func getPythPrice(assetID string) (float64, error) {
	apiURL := fmt.Sprintf("https://api.pyth.network/price/%s", assetID)

	resp, err := http.Get(apiURL)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var priceResponse PriceResponse
	err = json.Unmarshal(body, &priceResponse)
	if err != nil {
		return 0, err
	}

	return priceResponse.Price, nil
}

func main() {
	assetID := "SOLUSD" // Replace this with the desired asset ID
	price, err := getPythPrice(assetID)
	if err != nil {
		fmt.Printf("Error fetching price: %v\n", err)
		return
	}

	fmt.Printf("Price of %s: %f\n", assetID, price)

}
