package main
import "sort"
// array = [12,3,1,2,-6,5,-8,6]
// targetsum = 0
// output = [[-8,2,6],[-8,3,5],[-6,1,5]]
func ThreeNumberSum(array []int, target int) [][]int {
	// Write your code here.
    sort.Ints(array)
    triplets := [][]int{}

    for i:=0; i< len(array)-2; i++ {
        left, right := i+1, len(array)-1
        for left < right {
            cSum := array[i] + array[left] + array[right]
            if cSum == target {
                triplets = append(triplets, []int{array[i],array[left],array[right]})
                left = left +1
                right = right - 1
            } else if cSum < target {
                left = left + 1
            } else if cSum > target {
                right = right - 1
            }
        }
    }
	return triplets
}
