//go:build wireinject
// +build wireinject

package main

import (
	"calculator/calculator"

	"github.com/google/wire"
)

func InitializeCalculator() calculator.Calculator {
	wire.Build(
		calculator.AdderProvider,
		calculator.NewCalculator,
	)

	return calculator.Calculator{}
}
