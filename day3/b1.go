package day3

import (
	"adv-code-2023/common"
	"fmt"
	"unicode"
)

func B1(inputPath string) {
	// mapIndexSymbol:
	rows, cols, mapIndex, aNumbers := indexSymbol(inputPath)

	sum := int64(0)
	for _, aNumber := range aNumbers {
		if isInvalidNumber(rows, cols, aNumber, mapIndex) {
			num := aNumber.Number()
			sum += int64(num)
		}
	}
	fmt.Println("sum is", sum)
}

type ANumber struct {
	Symbols []rune
	Indexs  []int
}

func (a *ANumber) Size() int {
	return len(a.Symbols)
}

func (a *ANumber) Number() int {
	num := 0
	for idx, s := range a.Symbols {
		x := int(s - '0')
		for i := idx; i < len(a.Symbols)-1; i++ {
			x *= 10
		}
		num += x
	}
	return num
}

func indexSymbol(inputPath string) (int, int, map[int]struct{}, []ANumber) {
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
			if s != rune('.') {
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

func isInvalidNumber(rows, cols int, aNumber ANumber, mapIndex map[int]struct{}) bool {
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
	for idx := range idxNeedCheck {
		if idx < 0 {
			continue
		}
		if _, exist := mapIndex[idx]; exist {
			return true
		}
	}
	return false
}
