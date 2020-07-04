package dog

import (
	"fmt"
	"testing"
)

type TestData struct {
	data   int
	answer int
}

var testData = []TestData{
	{3, 21},
	{2, 14},
	{1, 7},
}

func TestYears(t *testing.T) {
	for _, test := range testData {
		ans := Years(test.data)
		if ans != test.answer {
			t.Error("Expectd", test.answer, "Got", ans)
		}
	}
}

func ExampleYears() {
	fmt.Println(Years(3))
	// Output:
	// 21
}

func ExampleYears_two() {
	// I think it trim spacs
	fmt.Println("     Hello world")
	// Output:
	// Hello world
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(7)
	}
}

func TestYearsTwo(t *testing.T) {
	for _, test := range testData {
		ans := YearsTwo(test.data)
		if ans != test.answer {
			t.Error("Expectd", test.answer, "Got", ans)
		}
	}
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(3))
	// Output:
	// 21
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(7)
	}
}
