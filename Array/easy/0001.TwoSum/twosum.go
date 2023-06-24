package twosum

// algorithm go in leetcode
// 0001 twoSum 两数之和
// Example: [intput] nums = [2, 7, 11, 15], target = 9
//  		[output] [0, 1]

func TwoSum(nums []int, target int) []int {
	// Initialize Dictionary: map[int]int ==> map[2]0
	sumIndexMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if _, ok := sumIndexMap[another]; ok {
			return []int{sumIndexMap[another], i}
		}
		// Add item from mapping, key: nums item, value: nums index
		sumIndexMap[nums[i]] = i
	}
	return nil
}
