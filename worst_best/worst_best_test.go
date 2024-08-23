package worstbest

import (
	"fmt"
	"testing"
)

const arrSize = 100000

func worstCaseIntSlice(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = i
	}
	return arr
}

func bestCaseIntSlice(size int) []int {
	arr := make([]int, size)
	mid := size / 2
	for i := 0; i < mid; i++ {
		arr[i] = i * 2
		arr[mid+i] = i*2 + 1
	}
	if size%2 != 0 {
		arr[mid] = mid
	}
	return arr
}

func BenchmarkQuickSortWorst(b *testing.B) {
	testName := fmt.Sprintf("worst case, pivot = arr[0], %d elements", arrSize)
	b.Run(testName, func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := worstCaseIntSlice(arrSize)
			b.StartTimer()
			quickSort(ar)
			b.StopTimer()
		}
	})

	testName = fmt.Sprintf("worst case, pivot median of three, %d elements", arrSize)
	b.Run(testName, func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := worstCaseIntSlice(arrSize)
			b.StartTimer()
			quickSortMedianPivot(ar)
			b.StopTimer()
		}
	})
}

func BenchmarkQuickSortBest(b *testing.B) {
	testName := fmt.Sprintf("best case, pivot = arr[0], %d elements", arrSize)
	b.Run(testName, func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := bestCaseIntSlice(arrSize)
			b.StartTimer()
			quickSort(ar)
			b.StopTimer()
		}
	})

	testName = fmt.Sprintf("best case, pivot median of three, %d elements", arrSize)
	b.Run(testName, func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := bestCaseIntSlice(arrSize)
			b.StartTimer()
			quickSortMedianPivot(ar)
			b.StopTimer()
		}
	})
}
