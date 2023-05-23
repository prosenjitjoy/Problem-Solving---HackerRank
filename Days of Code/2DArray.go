var max int32
flag := true
for i := 1; i < 5; i++ {
	for j := 1; j < 5; j++ {
		tmp := arr[i-1][j-1] + arr[i-1][j] + arr[i-1][j+1] + arr[i][j] + arr[i+1][j-1] + arr[i+1][j] + arr[i+1][j+1]
		if flag {
			max = tmp
			flag = false
		}
		if tmp > max {
			max = tmp
		}
	}
}

fmt.Println(max)