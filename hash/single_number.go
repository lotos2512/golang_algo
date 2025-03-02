package hash

func singleNumber1(nums []int) (minV int) {
	checkResult := make(map[int]int, len(nums)/2+1)
	for i := range nums {
		if _, ok := checkResult[nums[i]]; !ok {
			checkResult[nums[i]] = i
		} else {
			delete(checkResult, nums[i])
		}
	}
	for i := range checkResult {
		return i
	}
	return -1
}

func singleNumber2(nums []int) (minV int) {
	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}
