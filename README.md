# Benchmark Table
| Algorithm | n=10 | n=1000 | n=100000 |Memory costs
|-----------|------|--------|----------|-------------|
| bubbleSort | 3 | 6 | 11 | -
| selectionSort | 3 | 6 | 10 | -
| insertionSort | 3 | 6 | 10 | -
| mergeSort | 4 | 6 | 8 | +
| quickSort | 4 | 6 | 8 | +

# Quick Sort worst/best case (100000 elements)
| Pivot | worst case ns/op | best case ns/op  
|-------|------------------|----------------
| arr[0] | 99961083551 | 47691639317
| median of three | 44459556856 | 46396518851