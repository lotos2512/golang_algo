package binary_search

func searchMatrix(matrix [][]int, target int) bool {
	left, right := 0, len(matrix)-1
	for left <= right {
		mid := left + (right-left)/2
		if target >= matrix[mid][0] && target <= matrix[mid][len(matrix[mid])-1] {
			if res := search(matrix[mid], target); res != -1 {
				return true
			}
			return false
		} else if target > matrix[mid][0] {
			left = mid + 1
		} else if target < matrix[mid][len(matrix[mid])-1] {
			right = mid - 1
		}
	}
	return false
}

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	return -1
}
