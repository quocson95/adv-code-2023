package common

import (
	"bufio"
	"os"
)

func ReadLine(filePath string, cb func(lineStr string)) error {
	if cb == nil {
		return nil
	}
	readFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		cb(fileScanner.Text())
	}
	readFile.Close()
	return nil
}
