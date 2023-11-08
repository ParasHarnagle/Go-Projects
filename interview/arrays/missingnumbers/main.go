
//    You're given an unordered list of unique integers 
// nums  in the range [1,n] where n  represents the length of (nums + 2)
// This means that two numbers in this range are missing
//    from the list.
// 
//    Write a function that takes in this list and returns a new list with the two
//    missing numbers, sorted numerically.
// INput  = [1, 4, 3]
// Output = [2, 5] 

package main

func MissingNumbers(nums []int) []int {
    includedNums := map[int]bool{}
    for _,num := range nums {
        includedNums[num] = true
    }
    sol := make([]int, 0)
    for num := 1; num < len(nums) + 3; num ++ {
        if !includedNums[num] {
            sol = append(sol,num)
        }
    }
	return sol
}
    
