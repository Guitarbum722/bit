package main

import "time"

// Bitcoin represents the current value in USD
type Bitcoin struct {
	Bpi struct {
		USD struct {
			Code        string  `json:"code"`
			Description string  `json:"description"`
			Rate        string  `json:"rate"`
			RateFloat   float64 `json:"rate_float"`
		} `json:"USD"`
	} `json:"bpi"`
	Disclaimer string `json:"disclaimer"`
	Time       struct {
		Updated    string    `json:"updated"`
		UpdatedISO time.Time `json:"updatedISO"`
		Updateduk  string    `json:"updateduk"`
	} `json:"time"`
}

func main() {

}

func call() error {

	return nil
}
