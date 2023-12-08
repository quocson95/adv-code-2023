package day3

import (
	"adv-code-2023/common"
	"fmt"
	"unicode"
)

func B2(inputPath string) {
	// mapIndexSymbol:
	rows, cols, mapIndex, aNumbers := indexSymbol2(inputPath)

	sum := int64(0)
	validNumbers := make(map[int][]ANumber, 0)
	for _, aNumber := range aNumbers {
		idxsValid := isInvalidNumber2(rows, cols, aNumber, mapIndex)
		if len(idxsValid) == 0 {
			continue
		}
		for _, idx := range idxsValid {
			validNumbers[idx] = append(validNumbers[idx], aNumber)
		}
	}
	for _, aNaNumbers := range validNumbers {
		if len(aNaNumbers) == 0 {
			return
		}
		if len(aNaNumbers) != 2 {
			fmt.Printf("exception")
			continue
		}
		sum += int64(aNaNumbers[0].Number()) * int64(aNaNumbers[1].Number())
	}

	fmt.Println("sum is", sum)
}

func indexSymbol2(inputPath string) (int, int, map[int]struct{}, []ANumber) {
	index := 0
	mapIndexSymbol := make(map[int]struct{})
	aNumbers := make([]ANumber, 0)
	aNumber := ANumber{}
	cols, rows := 0, 0
	common.ReadLine(inputPath, func(lineStr string) {
		if cols == 0 {
			cols = len(lineStr)
		}
		for _, s := range lineStr {
			if s == rune('*') {
				mapIndexSymbol[index] = struct{}{}
			}
			if unicode.IsDigit(s) {
				aNumber.Symbols = append(aNumber.Symbols, s)
				aNumber.Indexs = append(aNumber.Indexs, index)
			} else {
				if aNumber.Size() > 0 {
					aNumbers = append(aNumbers, aNumber)
					aNumber = ANumber{}
				}
			}
			index++
		}
		rows++
	})
	return rows, cols, mapIndexSymbol, aNumbers
}

func isInvalidNumber2(rows, cols int, aNumber ANumber, mapIndex map[int]struct{}) []int {
	idxNeedCheck := make(map[int]struct{}, 0)
	for _, idx := range aNumber.Indexs {
		// idxNeedCheck = append(idxNeedCheck, idx-1, idx+1, idx-cols, idx+cols)
		idxNeedCheck[idx-1] = struct{}{}
		idxNeedCheck[idx+1] = struct{}{}
		idxNeedCheck[idx-cols] = struct{}{}
		idxNeedCheck[idx-cols-1] = struct{}{}
		idxNeedCheck[idx-cols+1] = struct{}{}
		idxNeedCheck[idx+cols] = struct{}{}
		idxNeedCheck[idx+cols-1] = struct{}{}
		idxNeedCheck[idx+cols+1] = struct{}{}
	}
	for _, idx := range aNumber.Indexs {
		delete(idxNeedCheck, idx)
	}
	idxValid := make([]int, 0)
	for idx := range idxNeedCheck {
		if idx < 0 {
			continue
		}
		if _, exist := mapIndex[idx]; exist {
			idxValid = append(idxValid, idx)
		}
	}
	return idxValid
}
