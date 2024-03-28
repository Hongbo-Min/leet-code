package main

func moveZeroes(nums []int) {
	leftPointer := 0
	for rightPointer := 0; rightPointer < len(nums); rightPointer++ {
		if nums[rightPointer] != 0 {
			nums[leftPointer], nums[rightPointer] = nums[rightPointer], nums[leftPointer]
			leftPointer++
		}
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
}
