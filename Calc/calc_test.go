package calc_test

import (
	"testing"

	calc "github.com/yaumianwar/calculator/Calc"
)

func TestCalculation(t *testing.T) {
	c := calc.Calc{}

	c.Data1 = float64(1.5)
	c.Data2 = float64(0.1)

	if int(c.AddNumber()) != 1 {
		t.Errorf("error")
	}
}

func TestSubstract(t *testing.T) {
	c := calc.Calc{}

	c.Data1 = float64(1.5)
	c.Data2 = float64(0.1)

	if int(c.SubstractNumber()) != 2 {
		t.Errorf("error")
	}
}
