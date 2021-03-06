package shape

import (
	"testing"

	"github.com/devetek/Go-Interface/helper"
)

var c = Circle{10}

func TestCircleArea(t *testing.T) {
	var testNumb float64 = 314.1592653589793

	aCircle := c.Area()

	if !helper.FloatEqualizer(aCircle, testNumb) {
		t.Errorf("TestCircleArea got %f want %f", aCircle, testNumb)
	}
}

func TestCirclePerimeter(t *testing.T) {
	var testNumb float64 = 32.3606797749979

	pCircle := c.Perimeter()

	if !helper.FloatEqualizer(pCircle, testNumb) {
		t.Errorf("TestCirclePerimeter got %f want %f", pCircle, testNumb)
	}
}
