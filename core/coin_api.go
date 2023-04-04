package core

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/siddhantprateek/go-gpt/config"
	model "github.com/siddhantprateek/go-gpt/model"
)

func LatestCoinData() {

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://sandbox-api.coinmarketcap.com/v1/cryptocurrency/listings/latest", nil)
	if err != nil {
		log.Fatal(err)
	}

	query := url.Values(nil)
	query.Add("start", "1")
	query.Add("limit", "5000")
	query.Add("convert", "USD")

	apiKey := config.GetCMCKey()

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", apiKey)
	req.URL.RawQuery = query.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request to server")
		os.Exit(1)
	}

	defer resp.Body.Close()

	var respData model.ResponseData
	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		log.Fatal(err)
	}

	// var data []model.Data
	fmt.Println(respData)

}
