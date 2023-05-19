package main

import (
	"fmt"
)

func Swap(a, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp
}

func BinarySearch(a []int, Left, Right, Key int) int {
	for Left < Right {
		Mid := (Left + Right) / 2
		if a[Mid] == Key {
			return 1
		}
		if a[Mid] > Key {
			Right = Mid - 1
		} else {
			Left = Mid + 1
		}
	}
	return -1

}

func InsertionSort(a []int) {
	for i := 0; i < len(a)-1; i++ {
		j := i
		Min := j
		for ; j < len(a); j++ {
			if a[Min] > a[j] {
				Min = j
			}
		}
		Swap(&a[Min], &a[i])
	}
}

func BubbleSort(a []int) {
	for i := 0; i < len(a); i++ {
		for j := len(a) - 1; j > i; j-- {
			if a[j] < a[j-1] {
				Swap(&a[j], &a[j-1])
			}
		}
	}
}

func QuickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	left, right := 0, len(a)-1

	pi := right

	for i, _ := range a {
		if a[i] < a[pi] {
			a[i], a[left] = a[left], a[i]
			left++
		}

	}
	a[left], a[pi] = a[pi], a[left]
	QuickSort(a[:left])
	QuickSort(a[left+1:])
	return a

}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	var i int = 5
	a = append(a[:i], a[i+1:]...)
	fmt.Print(a)
}
