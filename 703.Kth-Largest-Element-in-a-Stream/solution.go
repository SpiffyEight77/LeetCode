package leetcode

type KthLargest struct {
	nums []int
	k    int
}

func Constructor(k int, nums []int) KthLargest {

	obj := KthLargest{nums: []int{}, k: k}
	for _, v := range nums {
		obj.Add(v)
	}
	return obj
}

func (this *KthLargest) Add(val int) int {
	if len(this.nums) < this.k-1 {
		this.nums = append(this.nums, val)
		return -1
	} else if len(this.nums) == this.k-1 {
		this.nums = append(this.nums, val)
		for i := len(this.nums) / 2; i >= 0; i-- {
			sink(this.nums, i, len(this.nums))
		}
	} else {
		if val > this.nums[0] {
			this.nums[0] = val
			sink(this.nums, 0, len(this.nums))
		}
	}
	return this.nums[0]
}

func sink(nums []int, i, length int) {
	for {
		idx := i
		left, right := 2*i+1, 2*i+2

		if left < length && nums[left] < nums[idx] {
			idx = left
		}
		if right < length && nums[right] < nums[idx] {
			idx = right
		}
		if idx == i {
			break
		}

		swap(nums, i, idx)
		i = idx
	}
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
