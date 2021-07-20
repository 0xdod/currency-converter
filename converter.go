package main

import (
	"errors"
	"fmt"
	"strings"

	"strconv"
)

// Exchange rates uses USD as a reference currency.
var exchangeRates = map[string]float64{
	"NGN": 411.50,
	"KSH": 108.17,
	"GHS": 5.95,
}

type Currency struct {
	Name  string  `json:"currency"`
	Value float64 `json:"value"`
}

type ConversionResult struct {
	From Currency `json:"from"`
	To   Currency `json:"to"`
}

func Convert(from, to string, amount float64) (ConversionResult, error) {

	fromPerUSD := exchangeRates[strings.ToUpper(from)]
	toPerUSD := exchangeRates[strings.ToUpper(to)]

	if fromPerUSD == 0 || toPerUSD == 0 {
		return ConversionResult{}, errors.New("Not a valid currency")
	}

	value := (amount * toPerUSD) / fromPerUSD

	// Floating point arithmetic workaround
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)

	return ConversionResult{
		From: Currency{from, amount},
		To:   Currency{to, value},
	}, nil
}
