package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	m := map[string]string{}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	for i := 0; i < n; i++ {
		scanner.Scan()
		line = scanner.Text()
		line = strings.TrimSpace(line)
		input := strings.Split(line, " ")
		m[input[0]] = input[1]
	}

	for scanner.Scan() {
		name := scanner.Text()
		value, ok := m[name]
		if !ok {
			fmt.Println("Not found")
		} else {
			fmt.Printf("%s=%s\n", name, value)
		}
	}
}
