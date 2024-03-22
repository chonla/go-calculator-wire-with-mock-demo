package calculator_test

import (
	"calculator/calculator"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAdder struct {
	mock.Mock
}

func (m *MockAdder) Add(x, y int) int {
	args := m.Called(x, y)
	return args.Int(0)
}

func TestCalculatorAdderShouldCallAdderAdd(t *testing.T) {
	mockAdder := new(MockAdder)

	mockAdder.On("Add", 3, 6).Return(500)

	calc := calculator.NewCalculator(mockAdder)
	result := calc.Add(3, 6)

	mockAdder.AssertExpectations(t)
	assert.Equal(t, 500, result)
}
