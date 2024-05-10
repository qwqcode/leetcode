/**
 Do not return anything, modify nums1 in-place instead.
 */
function merge(nums1: number[], m: number, nums2: number[], n: number): void {
	const sorted: number[] = [];

	let p1 = 0;
	let p2 = 0;

	while (p1 < m || p2 < n) {
		if (p1 >= m) {
			sorted.push(...nums2.slice(p2))
			break
		}
		if (p2 >= n) {
			sorted.push(...nums1.slice(p1))
			break
		}

		if (nums1[p1] > nums2[p2]) {
			sorted.push(nums2[p2])
			p2++
		} else {
			sorted.push(nums1[p1])
			p1++
		}
	}

	for (let i = 0; i < nums1.length; i++) {
		nums1[i] = sorted[i]
	}
};