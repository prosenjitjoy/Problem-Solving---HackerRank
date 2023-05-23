bin := ""
for n != 0 {
	rem := n % 2
	n /= 2
	bin = string(rem+'0') + bin
}
groups := strings.Split(bin, "0")
max := -1
for _, group := range groups {
	if len(group) > max {
		max = len(group)
	}
}

fmt.Println(max)