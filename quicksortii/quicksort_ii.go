package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://www.hackerrank.com/challenges/quicksort2

func main() {
	next := func() func() int {
		scan := bufio.NewScanner(os.Stdin)
		scan.Split(bufio.ScanWords)
		return func() int {
			scan.Scan()
			i, _ := strconv.Atoi(scan.Text())
			return i
		}
	}()

	n := next()

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = next()
	}
	quicksort(arr)

}

func quicksort(arr []int) {
	if len(arr) <= 1 {
		return
	} else {
		pivot := arr[0]
		l := make([]int, 0)
		r := make([]int, 0)
		for _, k := range arr {
			if k < pivot {
				l = append(l, k)
			} else if k > pivot {
				r = append(r, k)
			}
		}
		quicksort(l)
		quicksort(r)

		sep := ""

		copy(arr, l)
		copy(arr[len(l)+1:], r)
		arr[len(l)] = pivot
		for _, i := range arr {
			fmt.Print(sep)
			fmt.Print(i)
			sep = " "
		}
		fmt.Println()
	}
}
