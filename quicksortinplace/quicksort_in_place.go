package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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
	for i := range arr {
		arr[i] = next()
	}
	quicksort(arr, 0, len(arr)-1)
}

func quicksort(arr []int, l, r int) {
	if r-l <= 0 {
		return
	} else {
		pivot := arr[r]
		i, j := l, l
		for j <= r-1 {
			if arr[j] <= pivot {
				tmp := arr[j]
				arr[j] = arr[i]
				arr[i] = tmp
				i++
			}
			j++
		}
		arr[r] = arr[i]
		arr[i] = pivot
		sep := ""
		for _, k := range arr {
			fmt.Print(sep)
			fmt.Print(k)
			sep = " "
		}
		fmt.Println()

		quicksort(arr, l, i-1)
		quicksort(arr, i+1, r)
	}
}
