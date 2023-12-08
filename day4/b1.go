package day4

import (
	"adv-code-2023/common"
	"fmt"
	"strconv"
	"strings"
)

func B1() {
	cardPoint := int64(0)
	common.ReadLine("day4/input_b1.txt", func(lineStr string) {
		lineStr = strings.Split(lineStr, ":")[1]
		winNum := make(map[int]struct{})
		arr := strings.Split(lineStr, "|")
		winNumStr := strings.TrimSpace(arr[0])
		for _, num := range ListStringToNumber(winNumStr) {
			winNum[num] = struct{}{}
		}
		myNums := ListStringToNumber(strings.TrimSpace(arr[1]))
		point := int64(0)
		for _, myNum := range myNums {
			if _, exist := winNum[myNum]; !exist {
				continue
			}
			if point == 0 {
				point = 1
				continue
			}
			point = point << 1
		}
		cardPoint += point
	})
	fmt.Println(cardPoint)
}

func ListStringToNumber(str string) []int {
	ml := make([]int, 0)
	for _, s := range strings.Split(str, " ") {
		str := strings.TrimSpace(string(s))
		if len(str) == 0 {
			continue
		}
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		ml = append(ml, num)
	}
	return ml
}
