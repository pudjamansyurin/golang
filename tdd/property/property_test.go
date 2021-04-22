package property

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	arabic int
	roman  string
}{
	{1, "I"},
	{2, "II"},
	{3, "III"},
	{4, "IV"},
	{5, "V"},
	{6, "VI"},
	{7, "VII"},
	{8, "VIII"},
	{9, "IX"},
	{10, "X"},
	{14, "XIV"},
	{18, "XVIII"},
	{20, "XX"},
	{39, "XXXIX"},
	{40, "XL"},
	{47, "XLVII"},
	{49, "XLIX"},
	{50, "L"},
	{100, "C"},
	{90, "XC"},
	{400, "CD"},
	{500, "D"},
	{900, "CM"},
	{1000, "M"},
	{1984, "MCMLXXXIV"},
	{3999, "MMMCMXCIX"},
	{2014, "MMXIV"},
	{1006, "MVI"},
	{798, "DCCXCVIII"},
}

func TestRomanNumerals(t *testing.T) {
	for _, tC := range testCases {
		desc := fmt.Sprintf(`%d converted to %s`, tC.arabic, tC.roman)
		t.Run(desc, func(t *testing.T) {
			got := ConvertToRoman(tC.arabic)
			want := tC.roman

			if got != want {
				t.Errorf(`got %q, want %q`, got, want)
			}
		})
	}
}

func TestRomanToArabic(t *testing.T) {
	for _, tC := range testCases {
		desc := fmt.Sprintf(`%s converted to %d`, tC.roman, tC.arabic)
		t.Run(desc, func(t *testing.T) {
			got := ConvertToArabic(tC.roman)
			want := tC.arabic

			if got != want {
				t.Errorf(`got %d, want %d`, got, want)
			}
		})
	}
}
