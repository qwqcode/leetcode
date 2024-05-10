function removeElement(nums: number[], val: number): number {
	let idx = 0;

	for (let i = 0; i < nums.length; i++) {
		if (nums[i] === val) continue;

		nums[idx] = nums[i];
		idx++;
	}

	return idx;
};