
//  A subsequence of an array is a set of numbers that aren't necessarily adjacent
//  in the array but that are in the same order as they appear in the array. For
//  instance, the numbers  form a subsequence of the array , and so do the numbers 
// [2, 4] .
//  Note
//  that a single number in an array and the array itself are both valid
//  subsequences of the array.
// array =  = [5, 1, 22, 25, 6, -1, 8, 10]
// subsequence =  = [1, 6, -1, 10]

package main

func IsValidSubsequence(array []int, sequence []int) bool {
	// Write your code here.
	var arrIdx, seqIdx int
	for arrIdx < len(array)  && seqIdx< len (sequence) {
		if array[arrIdx] == sequence[seqIdx] {
			seqIdx = seqIdx + 1
		}
		arrIdx = arrIdx + 1
	}
	return seqIdx == len(sequence)
}

  
