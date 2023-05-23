// Write your code here
numberOfSwaps := 0
for i := 0; i < int(n); i++ {
	for j := 0; j < int(n)-1; j++ {
		if a[j] > a[j+1] {
			a[j], a[j+1] = a[j+1], a[j]
			numberOfSwaps++
		}
	}
	if numberOfSwaps == 0 {
		break
	}
}
fmt.Printf("Array is sorted in %d swaps.\n", numberOfSwaps)
fmt.Println("First Element:", a[0])
fmt.Println("Last Element:", a[len(a)-1])