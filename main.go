/*
This is a silly command line tool and is not intended for serious use.
*/
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/pkg/errors"
)

// Bitcoin represents the current value in USD
type Bitcoin struct {
	Bpi struct {
		EUR struct {
			Code        string  `json:"code"`
			Description string  `json:"description"`
			Rate        string  `json:"rate"`
			RateFloat   float64 `json:"rate_float"`
			Symbol      string  `json:"symbol"`
		} `json:"EUR"`
		GBP struct {
			Code        string  `json:"code"`
			Description string  `json:"description"`
			Rate        string  `json:"rate"`
			RateFloat   float64 `json:"rate_float"`
			Symbol      string  `json:"symbol"`
		} `json:"GBP"`
		USD struct {
			Code        string  `json:"code"`
			Description string  `json:"description"`
			Rate        string  `json:"rate"`
			RateFloat   float64 `json:"rate_float"`
			Symbol      string  `json:"symbol"`
		} `json:"USD"`
	} `json:"bpi"`
	ChartName  string `json:"chartName"`
	Disclaimer string `json:"disclaimer"`
	Time       struct {
		Updated    string    `json:"updated"`
		UpdatedISO time.Time `json:"updatedISO"`
		Updateduk  string    `json:"updateduk"`
	} `json:"time"`
}

const usage = `Usage: bit [-c]
	-c       ISO 4217 Code (USD, EUR ...)

	USD will be used by default.
  `

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usage)
	}

	hFlag := flag.Bool("h", false, usage)
	helpFlag := flag.Bool("help", false, usage)
	cFlag := flag.String("c", "USD", "")
	flag.Parse()

	if *hFlag == true || *helpFlag == true {
		fmt.Fprintln(os.Stderr, errors.New(usage))
		os.Exit(0)
	}

	val := &Bitcoin{}

	err := call(*cFlag, val)
	if err != nil {
		fmt.Fprintln(os.Stderr, errors.Wrap(err, "err parsing json"))
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "Current Value in %v : %v  %v\n", *cFlag, val.Bpi.USD.Rate, val.Time.Updated)
}

func call(currency string, result interface{}) error {

	const endpoint = "https://api.coindesk.com/v1/bpi/currentprice/"
	uri := endpoint + currency + ".json"

	c := &http.Client{
		Timeout: time.Duration(time.Second * 20),
	}

	var req *http.Request
	var err error

	req, err = http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return err
	}
	defer func() { req.Close = true }()

	res, err := c.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(result)
}
