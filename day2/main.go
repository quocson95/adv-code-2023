package main

import (
	"adv-code-2023/common"
	"strconv"
	"strings"

	"go.uber.org/zap"
)

type BagColor struct {
	Red   int
	Green int
	Blue  int
}

func main() {
	logger, _ := zap.NewProduction()
	zap.ReplaceGlobals(logger)
	b2("input_b2.txt")
}

func b1(puzzleInputPath string) error {

	bagColorReq := BagColor{
		Red:   12,
		Green: 13,
		Blue:  14,
	}
	gameId := int64(0)
	sumGameId := int64(0)
	err := common.ReadLine(puzzleInputPath, func(lineStr string) {
		gameId++
		bagColor := parseGameColor(lineStr)
		if bagColor.Blue > bagColorReq.Blue {
			return
		}
		if bagColor.Red > bagColorReq.Red {
			return
		}
		if bagColor.Green > bagColorReq.Green {
			return
		}
		sumGameId += gameId
		// zap.L().With(zap.Any("bag", bagColor)).Info("test")

	})
	if err != nil {
		zap.L().With(zap.Error(err)).With(zap.String("file", puzzleInputPath)).Error("read file failed")
	}
	zap.L().With(zap.Int64("sun game id", sumGameId)).Info("done")
	return err
}

func b2(puzzleInputPath string) error {
	gameId := int64(0)
	power := int64(0)
	err := common.ReadLine(puzzleInputPath, func(lineStr string) {
		gameId++
		bagColor := parseGameColor(lineStr)
		// if bagColor.Blue > bagColorReq.Blue {
		// 	return
		// }
		// if bagColor.Red > bagColorReq.Red {
		// 	return
		// }
		// if bagColor.Green > bagColorReq.Green {
		// 	return
		// }
		power += int64(bagColor.Blue) * int64(bagColor.Red) * int64(bagColor.Green)
		// zap.L().With(zap.Any("bag", bagColor)).Info("test")

	})
	if err != nil {
		zap.L().With(zap.Error(err)).With(zap.String("file", puzzleInputPath)).Error("read file failed")
	}
	zap.L().With(zap.Int64("power", power)).Info("done")
	return err
}

func parseGameColor(str string) BagColor {
	bagColor := BagColor{}
	// remove Game x:
	str = strings.Split(str, ":")[1]
	// slit cubes each set
	sets := strings.Split(str, ";")
	for _, set := range sets {
		v := BagColor{}
		set = strings.TrimSpace(set)
		colors := strings.Split(set, ",")
		// parse num color
		for _, color := range colors {
			color = strings.TrimSpace(color)
			arr := strings.Split(color, " ")
			colorName := arr[1]
			colorNum, _ := strconv.Atoi(arr[0])
			if colorName == "red" {
				v.Red += colorNum
			}
			if colorName == "green" {
				v.Green += colorNum
			}
			if colorName == "blue" {
				v.Blue += colorNum
			}
		}
		if bagColor.Blue < v.Blue {
			bagColor.Blue = v.Blue
		}
		if bagColor.Red < v.Red {
			bagColor.Red = v.Red
		}
		if bagColor.Green < v.Green {
			bagColor.Green = v.Green
		}
	}
	return bagColor
}
