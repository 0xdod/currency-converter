package main

type CurrencyConverter struct {
	From   Currency
	To     Currency
	Amount float64
}

type Currency string

type ConversionResult struct {
	Currency
	Value float64
}

func Convert(from, to string, amount float64) ConversionResult {
	// conversion logic
	// check if currency is valid.

	return ConversionResult{
		Currency: Currency(to),
		Value:    0.55,
	}
}
