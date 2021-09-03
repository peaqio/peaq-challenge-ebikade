package utils

import (
	"fmt"
	"strconv"
)

func CalculatePercentageDifference(newNum, originNum float64) string {

	diff := newNum - originNum
	percentDiff := (diff / originNum) * 100

	return Round4Decimal(percentDiff)
}

func Round4Decimal(num float64) string {
	numb := fmt.Sprintf("%.4f", num)
	return numb
}

func ToFloat(value string) float64 {
	result, _ := strconv.ParseFloat(value, 64)
	return result
}
