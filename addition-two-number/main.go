package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Array das
type Array []interface{}

// Map dsa
func (arr Array) Map(cb func(val interface{}, index int) interface{}) []interface{} {
	results := make([]interface{}, len(arr))

	for index, val := range arr {
		val := cb(val, index)
		results = append(results, val)
	}

	return results
}

func addStringsAsInterger(numStrings []string) int {
	sum := 0
	arr := Array(numStrings)

	ints := arr.Map(func(val interface{}, index int) interface{} {
		val, err := strconv.Atoi(str)
		if err != nil {
			log.Fatalln(err)
		}
		return val
	})

	for _, val := range int {
		sum += val
	}

	return sum
}

func main() {
	scnr := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter number with spaces")

	for scnr.Scan() {
		line := scnr.Text()
		numStrings := strings.Fields(line)
		fmt.Println("Answer : ", addStringsAsInterger(numStrings))

	}
}
