package main

import (
	"fmt"
	"log"
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	Arabic int
	Roman  string
}{
	{Arabic: 1, Roman: "I"},
	{Arabic: 2, Roman: "II"},
	{Arabic: 3, Roman: "III"},
	{Arabic: 4, Roman: "IV"},
	{Arabic: 5, Roman: "V"},
	{Arabic: 6, Roman: "VI"},
	{Arabic: 7, Roman: "VII"},
	{Arabic: 8, Roman: "VIII"},
	{Arabic: 9, Roman: "IX"},
	{Arabic: 10, Roman: "X"},
	{Arabic: 14, Roman: "XIV"},
	{Arabic: 18, Roman: "XVIII"},
	{Arabic: 20, Roman: "XX"},
	{Arabic: 39, Roman: "XXXIX"},
	{Arabic: 40, Roman: "XL"},
	{Arabic: 47, Roman: "XLVII"},
	{Arabic: 49, Roman: "XLIX"},
	{Arabic: 50, Roman: "L"},
	{Arabic: 100, Roman: "C"},
	{Arabic: 90, Roman: "XC"},
	{Arabic: 400, Roman: "CD"},
	{Arabic: 500, Roman: "D"},
	{Arabic: 900, Roman: "CM"},
	{Arabic: 1000, Roman: "M"},
	{Arabic: 1984, Roman: "MCMLXXXIV"},
	{Arabic: 3999, Roman: "MMMCMXCIX"},
	{Arabic: 2014, Roman: "MMXIV"},
	{Arabic: 1006, Roman: "MVI"},
	{Arabic: 798, Roman: "DCCXCVIII"},
}

func TestConvertArabicToRomanNumerals(t *testing.T) {
	for _, kase := range cases {
		description := fmt.Sprintf("convert %d to %q", kase.Arabic, kase.Roman)

		t.Run(description, func(t *testing.T) {
			actual := ConvertToRoman(kase.Arabic)
			expected := kase.Roman

			assert.Equal(t, expected, actual)
		})
	}
}

func TestConvertRomanNumeralToArabic(t *testing.T) {
	for _, kase := range cases[:1] {
		description := fmt.Sprintf("convert %q to %d", kase.Roman, kase.Arabic)

		t.Run(description, func(t *testing.T) {
			actual := ConvertToArabic(kase.Roman)
			expected := kase.Arabic

			assert.Equal(t, expected, actual)
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			log.Println(arabic)
			return true
		}

		roman := ConvertToRoman(int(arabic))
		fromRoman := ConvertToArabic(roman)

		return fromRoman == int(arabic)
	}

	if err := quick.Check(assertion, &quick.Config{
		MaxCount: 1000,
	}); err != nil {
		t.Error("failed checks", err)
	}
}
