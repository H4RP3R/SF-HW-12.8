package sort

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
