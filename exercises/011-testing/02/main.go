package main

import (
	"fmt"
	"harish/try-go/exercises/011-testing/02/quote"
	"harish/try-go/exercises/011-testing/02/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}

	m := map[string]int{
		"a": 1,
	}

	fmt.Println(m["das"])
}
