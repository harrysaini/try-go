package mymath

import (
	"fmt"
	"testing"
)

type TestData struct {
	data   []int
	answer float64
}

var tests = []TestData{
	{[]int{1, 4, 6, 8, 100}, 6},
	{[]int{0, 8, 10, 1000}, 9},
	{[]int{9000, 4, 10, 8, 6, 12}, 9},
	{[]int{123, 744, 140, 200}, 170},
}

func TestCenteredAvg(t *testing.T) {
	for _, test := range tests {
		res := CenteredAvg(test.data)

		if res != test.answer {
			t.Error("expected", test.answer, "got", res)
		}
	}
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg(tests[0].data)
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{1, 4, 6, 8, 100}))
	// Output
	// 6
}
