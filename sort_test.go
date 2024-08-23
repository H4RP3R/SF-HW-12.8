package sort

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

const (
	testArrMaxSize = 150
	testArrMinVal  = -100
	testArrMaxVal  = 100
	testsCount     = 1000
)

func randomIntArr(size, min, max int) []int {
	max += 1
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(max-min) + min
	}
	return arr
}

func Test_bubbleSort(t *testing.T) {
	for i := 0; i < testsCount; i++ {
		size := rand.Intn(testArrMaxSize)
		min := testArrMinVal
		max := rand.Intn(testArrMaxVal)
		testName := fmt.Sprintf("arr size: %d, min: %d, max: %d", size, min, max)
		t.Run(testName, func(t *testing.T) {
			testArr := randomIntArr(size, min, max)
			gotArr := make([]int, size)
			copy(gotArr[:], testArr[:])
			sort.Ints(testArr)
			bubbleSort(gotArr)
			if !reflect.DeepEqual(testArr, gotArr) {
				t.Errorf("slices aren't equal")
			}
		})
	}
}

func Test_selectionSort(t *testing.T) {
	for i := 0; i < testsCount; i++ {
		size := rand.Intn(testArrMaxSize)
		min := testArrMinVal
		max := rand.Intn(testArrMaxVal)
		testName := fmt.Sprintf("arr size: %d, min: %d, max: %d", size, min, max)
		t.Run(testName, func(t *testing.T) {
			testArr := randomIntArr(size, min, max)
			gotArr := make([]int, size)
			copy(gotArr[:], testArr[:])
			sort.Ints(testArr)
			selectionSort(gotArr)
			if !reflect.DeepEqual(testArr, gotArr) {
				t.Errorf("slices aren't equal")
			}
		})
	}
}

func Test_insertionSort(t *testing.T) {
	for i := 0; i < testsCount; i++ {
		size := rand.Intn(testArrMaxSize)
		min := testArrMinVal
		max := rand.Intn(testArrMaxVal)
		testName := fmt.Sprintf("arr size: %d, min: %d, max: %d", size, min, max)
		t.Run(testName, func(t *testing.T) {
			testArr := randomIntArr(size, min, max)
			gotArr := make([]int, size)
			copy(gotArr[:], testArr[:])
			sort.Ints(testArr)
			insertionSort(gotArr)
			if !reflect.DeepEqual(testArr, gotArr) {
				t.Errorf("slices aren't equal")
			}
		})
	}
}

func Test_mergeSort(t *testing.T) {
	for i := 0; i < testsCount; i++ {
		size := rand.Intn(testArrMaxSize)
		min := testArrMinVal
		max := rand.Intn(testArrMaxVal)
		testName := fmt.Sprintf("arr size: %d, min: %d, max: %d", size, min, max)
		t.Run(testName, func(t *testing.T) {
			testArr := randomIntArr(size, min, max)
			gotArr := make([]int, size)
			copy(gotArr[:], testArr[:])
			sort.Ints(testArr)
			gotArr = mergeSort(gotArr)
			if !reflect.DeepEqual(testArr, gotArr) {
				t.Errorf("slices aren't equal")
			}
		})
	}
}

func Test_quickSort(t *testing.T) {
	for i := 0; i < testsCount; i++ {
		size := rand.Intn(testArrMaxSize)
		min := testArrMinVal
		max := rand.Intn(testArrMaxVal)
		testName := fmt.Sprintf("arr size: %d, min: %d, max: %d", size, min, max)
		t.Run(testName, func(t *testing.T) {
			testArr := randomIntArr(size, min, max)
			gotArr := make([]int, size)
			copy(gotArr[:], testArr[:])
			sort.Ints(testArr)
			gotArr = quickSort(gotArr)
			if !reflect.DeepEqual(testArr, gotArr) {
				t.Errorf("slices aren't equal")
			}
		})
	}
}
