package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://www.hackerrank.com/challenges/insertionsort1

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

	picked := arr[len(arr)-1]
	for i := len(arr) - 2; i >= 0; i-- {
		if picked < arr[i] {
			arr[i+1] = arr[i]
			printSlice(arr)
			if i == 0 {
				arr[0] = picked
			}
		} else {
			arr[i+1] = picked
			break
		}
	}
	printSlice(arr)
}

func printSlice(s []int) {
	for _, i := range s {
		fmt.Print(i, " ")
	}
	fmt.Println()
}
