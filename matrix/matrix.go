package common

func transposition(items [][]int) [][]int {

	h := len(items[0])
	w := len(items)
	if h != w {
		newMatrix := make([][]int, h, w)
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				newMatrix[i] = append(newMatrix[i], items[j][i])
			}
		}
		items = newMatrix
	} else {
		for i := 0; i < h; i++ {
			for j := i + 1; j < w; j++ {
				v := items[i][j]
				items[i][j] = items[j][i]
				items[j][i] = v
			}
		}
	}
	return items
}
