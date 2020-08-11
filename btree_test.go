package main

import (
	"math/rand"
	"testing"
)

// func TestSelectionSort(t *testing.T) {
// 	got := generateList(1e2)
// 	want := make([]int, 1e2)
// 	copy(want, got)

// 	sort.Ints(want)
// 	selectionSort(got)

// 	if !reflect.DeepEqual(want, got) {
// 		t.Fatalf("expected: %v, got: %v", want, got)
// 	}
// }

func BenchmarkConstructTree100(b *testing.B) {
	b.ReportAllocs()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		numbers := generateList(1e3)
		b.StartTimer()
		var tree Btree

		for _, value := range numbers {
			tree.Insert(value)
		}
	}
}

func BenchmarkInsert(b *testing.B) {
	b.ReportAllocs()

	numbers := generateList(200)
	var tree Btree
	for i := 0; i < 100; i++ {
		tree.Insert(numbers[i])
	}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		b.StopTimer()
		value := rand.Intn(200)
		b.StartTimer()
		tree.Insert(value)
	}
}

func generateList(totalNumbers int) []int {
	numbers := make([]int, totalNumbers)
	for i := 0; i < totalNumbers; i++ {
		numbers[i] = rand.Intn(totalNumbers)
	}
	return numbers
}
