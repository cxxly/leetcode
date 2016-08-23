**heapsort**
```go
//max heap
func adjust(nums []int, root int, size int) {
	for {
		child := root<<1 + 1
		if child >= size {
			break
		}
		if child+1 < size && nums[child] < nums[child+1] {
			child++
		}
		if nums[root] >= nums[child] {
			return
		}
		nums[root], nums[child] = nums[child], nums[root]
		root = child
	}
}

func heapSort(nums []int) {
	heapsize := len(nums)
	for i := (len(nums) - 1) >> 1; i >= 0; i-- {
		adjust(nums, i, heapsize)
	}
	for i := len(nums) - 1; i >= 1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapsize -= 1
		adjust(nums, 0, heapsize)
	}
}
```
