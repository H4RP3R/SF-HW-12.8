package worstbest

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

func quickSortMedianPivot(ar []int) []int {
	if len(ar) < 2 {
		return ar
	}

	less := []int{}
	greater := []int{}
	equal := []int{}
	pivot := median(ar[0], ar[len(ar)/2], ar[len(ar)-1])
	for _, n := range ar {
		if n < pivot {
			less = append(less, n)
		} else if n > pivot {
			greater = append(greater, n)
		} else {
			equal = append(equal, n)
		}
	}

	return append(append(quickSort(less), equal...), quickSort(greater)...)
}

func median(a, b, c int) int {
	if a <= b && b <= c {
		return b
	}
	if a <= c && c <= b {
		return c
	}
	if b <= a && a <= c {
		return a
	}
	if b <= c && c <= a {
		return c
	}
	if c <= a && a <= b {
		return a
	}
	return b
}
