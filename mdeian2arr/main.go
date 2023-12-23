package main

import "sort"


func findmed(arr1. arr2 []int) float64  {
	// combine 2 arrays
	c := append(arr1,arr2...)
	// sort the slice
	sort.Ints(c)
	// cal the median
	median := len(c)/2
	if len(c)%2 == 0 {
		// even number of elements average the middle two
		return float64(c[med-1]+c[med])/2
	}
	if odd no. of elements
	return float64(c[med])

}