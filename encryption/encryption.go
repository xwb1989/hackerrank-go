package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Split(bufio.ScanLines)

	scan.Scan()
	org := scan.Text()

	row := int(math.Sqrt(float64(len(org))))
	col := int(math.Ceil(math.Sqrt(float64(len(org)))))

	if row*col < len(org) {
		row++
	}

	sep := ""
	for i := 0; i < col; i++ {
		fmt.Println("i: ", i)
		res := make([]byte, 0, col)
		for j := 0; j < row; j++ {
			idx := i + j*col
			if idx < len(org) {
				res = append(res, org[idx])
			} else {
				break
			}
		}
		fmt.Print(sep, string(res))
		sep = " "
	}
	fmt.Println()
}
