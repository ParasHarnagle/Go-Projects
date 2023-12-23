package main

func longestIncreasingSub(arr []int) []int {
	n := len(arr)
	if n == 0 {
		return []int{}
	}

	// initalize of lis
	lis := make([]int, n)
	for i := range lis {
		lis[i] = 1
	}

	// calculate the lis for each element
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] && lis[i] < lis[j]+1 {
				lis[i] = lis[j] + 1
			}
		}
	}
	// reconstruct the lis
	maxLength := 0
	for _, length := range lis {
		if length > maxLength {
			maxLength = length
		}
	}
	result := make([]int)

	return nil
}
