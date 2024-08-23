package sort

// Bubble sort
func bubbleSort(ar []int) {
	sorted := false

	for !sorted {
		sorted = true
		for i := 0; i < len(ar)-1; i++ {
			if ar[i] > ar[i+1] {
				ar[i], ar[i+1] = ar[i+1], ar[i]
				sorted = false
			}
		}
	}
}

// Selection sort
func selectionSort(ar []int) {
	if len(ar) < 2 {
		return
	}

	for i := 0; i < len(ar); i++ {
		minIdx := i
		for j := i + 1; j < len(ar); j++ {
			if ar[j] < ar[minIdx] {
				minIdx = j
			}
		}
		ar[i], ar[minIdx] = ar[minIdx], ar[i]
	}
}

// Insertion sort
func insertionSort(ar []int) {
	for i := 1; i < len(ar); i++ {
		for j := i; j > 0; j-- {
			if ar[j] < ar[j-1] {
				ar[j], ar[j-1] = ar[j-1], ar[j]
			} else {
				break
			}
		}
	}
}

// Merge sort
func mergeSort(ar []int) []int {
	if len(ar) < 2 {
		return ar
	}

	mid := len(ar) / 2
	left := mergeSort(ar[:mid])
	right := mergeSort(ar[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	merged := []int{}
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			merged = append(merged, left[i])
			i++
		} else {
			merged = append(merged, right[j])
			j++
		}
	}

	merged = append(merged, left[i:]...)
	merged = append(merged, right[j:]...)

	return merged
}

// Quick sort
func quickSort(ar []int) []int {
	if len(ar) < 2 {
		return ar
	}

	less := []int{}
	greater := []int{}
	pivot := ar[0]
	for _, n := range ar[1:] {
		if n < pivot {
			less = append(less, n)
		} else {
			greater = append(greater, n)
		}
	}

	return append(append(quickSort(less), pivot), quickSort(greater)...)
}
