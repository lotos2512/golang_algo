package find_anagrams

func findAnagrams(s string, p string) (result []int) {
	search := []rune(s)
	source := []rune(p)

	find := 0
	check := 0
	for i := 0; i < len(search); {
		for j := len(source) - 1; j >= 0; j-- {
			if source[j] == search[i] {
				find++
				break
			}
		}
		check++
		i++

		if check == len(source) {
			if find == len(source) {
				result = append(result, i-len(source))
			}
			check = 0
			find = 0
		}

	}
	return result
}
