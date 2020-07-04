package utils

import (
	"reflect"
	"testing"
)

func TestIsMapEqual(t *testing.T) {
	m1 := map[string]int{
		"a": 1,
		"b": 2,
	}
	m2 := map[string]int{
		"a": 1,
		"b": 2,
	}

	if !IsMapEqual(m1, m2) {
		t.Error("Expected", true, "Got", false)
	}
}

func TestIsMapEqual_two(t *testing.T) {
	m1 := map[string]int{
		"a": 1,
		"b": 2,
	}
	m2 := map[string]int{
		"a": 1,
	}

	if IsMapEqual(m1, m2) {
		t.Error("Expected", false, "Got", true)
	}
}

func TestIsMapEqual_three(t *testing.T) {
	m1 := map[string]int{
		"a": 1,
		"b": 2,
	}
	m2 := map[string]int{
		"a": 1,
		"b": 24,
	}

	if IsMapEqual(m1, m2) {
		t.Error("Expected", false, "Got", true)
	}
}

func BenchmarkIsMapEqual(b *testing.B) {
	m1 := map[string]int{
		"a": 1,
		"b": 2,
	}
	m2 := map[string]int{
		"a": 1,
		"b": 2,
	}

	for i := 0; i < b.N; i++ {
		IsMapEqual(m1, m2)
	}
}

func BenchmarkIsMapEqual_two(b *testing.B) {
	m1 := map[string]int{
		"a": 1,
		"b": 2,
	}
	m2 := map[string]int{
		"a": 1,
		"b": 2,
	}

	for i := 0; i < b.N; i++ {
		reflect.DeepEqual(m1, m2)
	}
}
