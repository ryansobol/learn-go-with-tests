package main

import "strings"

type RomanNumeral struct {
	Value  int
	Symbol string
}

var romanNumerals = []RomanNumeral{
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
	var result strings.Builder

	for _, romanNumeral := range romanNumerals {
		for arabic >= romanNumeral.Value {
			result.WriteString(romanNumeral.Symbol)

			arabic -= romanNumeral.Value
		}
	}

	return result.String()
}

func ConvertToArabic(roman string) int {
	arabic := 0

	for _, romanNumeral := range romanNumerals {
		for strings.HasPrefix(roman, romanNumeral.Symbol) {
			arabic += romanNumeral.Value

			roman = strings.TrimPrefix(roman, romanNumeral.Symbol)
		}
	}

	return arabic
}
