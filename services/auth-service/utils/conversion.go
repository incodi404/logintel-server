package utils

import "strconv"

type ConversionStru struct{}

func (c *ConversionStru) IntToStr(value int) string {
	return strconv.Itoa(value)
}

func (c *ConversionStru) StrToInt(value string) int {
	val, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}

	return val
}

var Conversion ConversionStru
