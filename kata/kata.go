package main

import (
	"strings"
)

type RomanNumeral struct {
	Value  uint16
	Symbol string
}

type RomanNumerals []RomanNumeral

func (r RomanNumerals) ValueOf(symbol string) uint16 {
	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}

	return 0
}

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

func ConvertToArabic(roman string) uint16 {
	var total uint16 = 0
	for i := 0; i < len(roman); i++ {

		if i+1 < len(roman) {
			potentialSymbol := string([]byte{roman[i], roman[i+1]})
			potentialValue := allRomanNumerals.ValueOf(potentialSymbol)
			if potentialValue != 0 {
				total += potentialValue
				i++
			} else {
				total += allRomanNumerals.ValueOf(string(roman[i]))
			}
		} else {
			total += allRomanNumerals.ValueOf(string(roman[i]))
		}
	}

	return total
}

func ConvertToRoman(arabic uint16) string {
	var result strings.Builder
	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String()
}
