package main

import (
	"adv-code-2023/day3"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	zap.ReplaceGlobals(logger)
	day3.B2("day3/input_b1.txt")
}
