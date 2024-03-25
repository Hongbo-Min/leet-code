package main

import "fmt"

func removeDuplicates(nums []int) int {
    leftPointer := 0
    for rightPointer := leftPointer + 1; rightPointer < len(nums); rightPointer++ {
        if nums[leftPointer] != nums[rightPointer] {
            leftPointer++
            nums[leftPointer] = nums[rightPointer]
        }
    }
    return leftPointer+1
}

// 删除有序数组中的重复项
func main() {
    nums := []int{0,0,1,1,1,2,2,3,3,4}
    fmt.Println(removeDuplicates(nums))
}
