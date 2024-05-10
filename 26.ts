function removeDuplicates(nums: number[]): number {
	let fast = 1
	let slow = 1

	for (; fast < nums.length; fast++) {
		if (nums[fast] != nums[fast-1]) {
			nums[slow] = nums[fast]
			slow++
		}
	}

	return slow
};