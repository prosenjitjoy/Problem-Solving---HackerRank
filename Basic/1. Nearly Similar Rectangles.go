/*
 * Complete the 'nearlySimilarRectangles' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts 2D_LONG_INTEGER_ARRAY sides as parameter.
 */

func nearlySimilarRectangles(sides [][]int64) int64 {
	var count int64 = 0
	m := map[float64]int64{}
	for _, side := range sides {
		m[float64(side[0])/float64(side[1])]++
	}

	for _, val := range m {
		if val > 1 {
			count += (val * (val - 1)) / 2
		}
	}

	return count
}