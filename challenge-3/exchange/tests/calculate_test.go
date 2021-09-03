package test

import (
	"fmt"
	"testing"

	"github.com/ebikode/peaq-challenge/challenge-3/exchange/utils"
)

// TestMain ...
func TestCalculate(t *testing.T) {

	originNum := 0.00000857
	newNum := 0.00000740

	got := utils.CalculatePercentageDifference(newNum, originNum)

	want := utils.Round4Decimal(-13.652275379229861)

	if got != want {
		err := fmt.Sprintf(`got "%s", want "%s"`, got, want)
		t.Fatal(err)
	}
}

func TestRound4Decimal(t *testing.T) {

	num := -13.652275379229861

	got := utils.Round4Decimal(num)

	want := "-13.6523"

	if got != want {
		err := fmt.Sprintf(`got "%s", want "%s"`, got, want)
		t.Fatal(err)
	}
}

func TestToFloat(t *testing.T) {

	num := "-13.652275379229861"

	got := utils.ToFloat(num)

	want := -13.652275379229861

	if got != want {
		err := fmt.Sprintf(`got "%f", want "%f"`, got, want)
		t.Fatal(err)
	}
}
