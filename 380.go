type RandomizedSet struct {
	m map[int]int // val => pos
	nums []int
}


func Constructor() RandomizedSet {
	instance := RandomizedSet{
		m: map[int]int{},
		nums: []int{},
	}
	return instance
}


func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; ok {
		return false
	}

	this.m[val] = len(this.nums) // record the position of the new value
	this.nums = append(this.nums, val) // append the value

	return true
}


func (this *RandomizedSet) Remove(val int) bool {
	id, ok := this.m[val] 
	if !ok {
		return false
	}

	last := len(this.nums) - 1

	// replace remove target value to last value
	this.nums[id] = this.nums[last]

	// update map last value pos to remove target pos
	this.m[this.nums[last]] = this.m[val]

	// delete remove target pos in map
	delete(this.m, val)
	
	// reduce the length of nums slice
	this.nums = this.nums[:len(this.nums) - 1]

	return true
}


func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */