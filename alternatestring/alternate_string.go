package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 100001)
	line, _, _ := reader.ReadLine()
	t, _ := strconv.Atoi(string(line))

	for ; t > 0; t-- {
		text, _, _ := reader.ReadLine()
		cnt := 0
		for i := 0; i < len(text)-1; i++ {
			if text[i] == text[i+1] {
				cnt++
			}
		}
		fmt.Println(cnt)
	}

}
