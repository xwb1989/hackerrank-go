package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://www.hackerrank.com/challenges/quicksort1

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
	l := next()
	arr := make([]int, l)
	for i := 0; i < l; i++ {
		arr[i] = next()
	}

	pivot := arr[0]
	i, j := l-1, l-1
	for j > 0 {
		if arr[j] >= pivot {
			tmp := arr[j]
			arr[j] = arr[i]
			arr[i] = tmp
			i--
		}
		j--
	}
	arr[0] = arr[i]
	arr[i] = pivot
	for _, i := range arr {
		fmt.Print(i)
		fmt.Print(" ")
	}
	fmt.Println()
}
