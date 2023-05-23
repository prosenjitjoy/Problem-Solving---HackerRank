package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var s string
		fmt.Scan(&s)
		a := ""
		b := ""
		for i, v := range s {
			if i%2 == 0 {
				a += string(v)
			} else {
				b += string(v)
			}
		}
		fmt.Println(a, b)
	}
}
