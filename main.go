package main

import (
	"adv-code-2023/day4"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	zap.ReplaceGlobals(logger)
	day4.B1()
}
