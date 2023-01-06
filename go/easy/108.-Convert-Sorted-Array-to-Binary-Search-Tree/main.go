package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return travel(nums, 0, len(nums)-1)
}

func travel(nums []int, lo, hi int) *TreeNode {
	if lo > hi {
		return nil
	}

	mid := lo + (hi-lo)/2

	return &TreeNode{
		Val:   nums[mid],
		Left:  travel(nums, lo, mid-1),
		Right: travel(nums, mid+1, hi),
	}
}
