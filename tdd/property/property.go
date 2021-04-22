package property

import "strings"

type RomanNumeral struct {
	Value  int
	Symbol string
}

type RomanNumerals []RomanNumeral

var allRomanNumerals = RomanNumerals{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabic int) string {
	var res strings.Builder

	for _, v := range allRomanNumerals {
		for arabic >= v.Value {
			res.WriteString(v.Symbol)
			arabic -= v.Value
		}
	}

	return res.String()
}

func (r RomanNumerals) ValueOf(symbols ...byte) int {
	symbol := string(symbols)
	for _, v := range r {
		if v.Symbol == symbol {
			return v.Value
		}
	}
	return 0
}

func ConvertToArabic(roman string) int {
	total := 0

	for i := 0; i < len(roman); i++ {
		symbol := roman[i]

		// look ahead
		if couldBeSubtracted(i, symbol, roman) {
			value := allRomanNumerals.ValueOf(symbol, roman[i+1])
			if value != 0 {
				total += value
				i++
			} else {
				total += allRomanNumerals.ValueOf(symbol)
			}
		} else {
			total += allRomanNumerals.ValueOf(symbol)
		}
	}
	return total
}

func couldBeSubtracted(i int, symbol byte, roman string) bool {
	isSubstractiveSymbol := symbol == 'I' || symbol == 'X' || symbol == 'C'
	return i+1 < len(roman) && isSubstractiveSymbol
}
