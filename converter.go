package main

import (
	"fmt"
	"strings"
	"time"

	"strconv"
)

// Exchange rates uses USD as a reference currency.
var exchangeRates = map[string]float64{
	"NGN": 411.50,
	"KSH": 108.17,
	"GHS": 5.95,
}

type Currency struct {
	Name  string
	Value float64
}

type ConversionResult struct {
	From Currency
	To   Currency
	Time time.Time
}

func Convert(from, to string, amount float64) ConversionResult {
	// conversion logic
	// check if currency is valid.

	fromPerUSD := exchangeRates[strings.ToUpper(from)]
	toPerUSD := exchangeRates[strings.ToUpper(to)]

	value := (amount * toPerUSD) / fromPerUSD

	// Floating point arithmetic workaround
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)

	return ConversionResult{
		From: Currency{from, amount},
		To:   Currency{to, value},
		Time: time.Now(),
	}
}
