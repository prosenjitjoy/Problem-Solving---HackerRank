/*
 * Complete the 'stringAnagram' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. STRING_ARRAY dictionary
 *  2. STRING_ARRAY query
 */

func stringAnagram(dictionary []string, query []string) []int32 {
	ans := make([]int32, len(query))

	m := map[[26]int]int32{}
	for _, val := range dictionary {
		var temp [26]int
		for i := 0; i < len(val); i++ {
			temp[val[i]-'a']++
		}
		m[temp]++
	}

	for idx, val := range query {
		var temp [26]int
		for i := 0; i < len(val); i++ {
			temp[val[i]-'a']++
		}
		ans[idx] = m[temp]
	}

	return ans
}