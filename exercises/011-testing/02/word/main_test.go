package word

import (
	"fmt"
	"harish/try-go/exercises/011-testing/02/utils"
	"testing"
)

func TestUseCount(t *testing.T) {
	expected := map[string]int{
		"I":      1,
		"am":     1,
		"Groot,": 1,
		"rocks":  2,
	}

	testString := "I am Groot, rocks rocks"

	result := UseCount(testString)

	if ok := utils.IsMapEqual(result, expected); !ok {
		t.Error("Expected", expected, "Got", result)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount("I am Groot, rocks rocks")
	}
}

func TestCount(t *testing.T) {
	testString := "I am Groot, rocks rocks"
	expected := 5
	result := Count(testString)

	if result != expected {
		t.Error("Expected", expected, "Got", result)
	}
}

func ExampleCount() {
	testString := "I am Groot, rocks rocks"
	result := Count(testString)
	fmt.Println(result)
	// Output:
	// 5
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count("I am Groot, rocks rocks")
	}
}
