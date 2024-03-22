package calculator_test

import (
	"calculator/calculator"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddShouldSumValueOfParameters(t *testing.T) {
	adder := calculator.NewAdder()

	result := adder.Add(1, 2)

	assert.Equal(t, 3, result)
}
