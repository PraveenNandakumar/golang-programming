/*

Problem: Two Sum - Given an array of integer numbers and an integer target,
return indices of the two numbers such that they add up to target.

Source: https://leetcode.com/problems/two-sum/

Example 1:
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [1,0].

*/

package main

import "fmt"

func twoSum(nums []int, target int) []int {
	// Create a hash map to store the complement of each number in the array and its index.
	complementMap := make(map[int]int)

	// Iterate through the array
	for i, num := range nums {
		// Check if complement of the number exists in the map.
		if complementIndex, ok := complementMap[num]; ok {
			// If exists, return the current index and complement index.
			return []int{i, complementIndex}
		}

		//Otherwise, add the complement of current number to the map
		complement := target - num
		complementMap[complement] = i
	}

	// If no numbers add upto the target, return empty array.
	return []int{}

}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	result := twoSum(nums, target)
	fmt.Println(result)
}

// Time Complexity: O(n)
// Space Complexity: O(n)
