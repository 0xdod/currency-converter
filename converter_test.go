package main

import (
	"fmt"
	"testing"
)

func TestConvert(t *testing.T) {

	cases := []struct {
		ConvertFrom string
		ConvertTo   string
		Value       float64
		Expected    float64
		expectError bool
	}{
		{"ksh", "ghs", 20, 1.10, false},
		{"ghs", "ksh", 30, 545.39, false},
		{"ksh", "ngn", 100, 380.42, false},
		{"ngn", "ksh", 100, 26.29, false},
		{"ghs", "ngn", 10, 691.60, false},
		{"ngn", "ghs", 1000, 14.46, false},
		{"gbh", "xyz", 100, 0.00, true},
		{"", "", 100, 0.00, true},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("Convert from %q to %q", tc.ConvertFrom, tc.ConvertTo), func(t *testing.T) {
			res, err := Convert(tc.ConvertFrom, tc.ConvertTo, tc.Value)

			if tc.expectError {
				if err == nil {
					t.Error("Expected an error, but got none")
				}
			} else {
				got := res.To.Value
				want := tc.Expected

				if got != want {
					t.Errorf("got %.2f, want %.2f", got, want)
				}
			}

		})
	}

}
