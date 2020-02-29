package codestyle

import "fmt"

func search(n int, callback func(i int) bool) int {
	l, r := 0, n
	for l < r {
		m := int(uint(l+r) >> 1)
		if callback(m) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func BinarySearch(a []int, t int) {
	i := search(len(a), func(i int) bool {
		return a[i] >= t
	})
	if i < len(a) && a[i] == t {
		fmt.Println(i)
	} else {
		fmt.Printf(`// %v is not present in data,
// but %v is the index where it would be inserted.
`, t, i)
	}
}

func BinarySearchUsage() {
	BinarySearch([]int{2, 3, 4, 5, 6}, 3)
	BinarySearch([]int{2, 3, 4, 5, 6}, -1)
	BinarySearch([]int{2, 3, 4, 5, 6}, 9)
}
