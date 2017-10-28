package main

import (
	"io/ioutil"
	"fmt"
	"os"
	"encoding/json"
	"github.com/toorop/go-bittrex"
	"github.com/ashwanthkumar/slack-go-webhook"
	"path/filepath"
)

var c Config
var bt *bittrex.Bittrex
var usdPerBtc float64

func main() {
	c = getConfig()

	bt = bittrex.New(c.Credentials.ApiKey, c.Credentials.ApiSecret)

	getBtcPerUsd()
	processPairs(c.Pairs)
}

func getConfig() Config {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	raw, err := ioutil.ReadFile(dir + "/config.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c Config
	json.Unmarshal(raw, &c)
	return c
}

func getBtcPerUsd() {
	usdTicker, err := bt.GetTicker("USDT-BTC")
	if err != nil {
		panic("Can't get USDT-BTC pair.")
	}

	usdPerBtc = usdTicker.Bid
}

func processPairs(pairs Pairs) {
	for _, pair := range pairs {
		processPair(pair)
	}
}

func processPair(pair Pair) {
	ticker := pairTicker(pair)
	tickerPriceInUsd := btcToUsd(ticker.Bid)

	if pair.Above && tickerPriceInUsd >= pair.Threshold {
		alertPriceAboveThreshold(pair.Symbol, tickerPriceInUsd, pair.Threshold)
	}

	if !pair.Above && tickerPriceInUsd <= pair.Threshold {
		alertPriceBelowThreshold(pair.Symbol, tickerPriceInUsd, pair.Threshold)
	}
}

func btcToUsd(btc float64) float64 {
	return btc * usdPerBtc
}

func pairTicker(pair Pair) bittrex.Ticker {
	ticker, err := bt.GetTicker(pair.Symbol)
	if err != nil {
		panic("Can't get " + pair.Symbol + " pair.")
	}

	return ticker
}

func alertPriceAboveThreshold(symbol string, price, threshold float64) {
	message := fmt.Sprintf("The price of %s pair ($%.2f) is above the threshold ($%.2f).", symbol, price, threshold)
	sendSlackMessage(message)
}

func alertPriceBelowThreshold(symbol string, price, threshold float64) {
	message := fmt.Sprintf("The price of %s pair ($%.2f) is below the threshold ($%.2f).", symbol, price, threshold)
	sendSlackMessage(message)
}

func sendSlackMessage(message string) {
	payload := slack.Payload{
		Text:      message,
		Username:  c.Slack.UserName,
		IconEmoji: c.Slack.IconEmoji,
	}
	slack.Send(c.Slack.WebHookUrl, "", payload)
}
