package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	t, _ := strconv.Atoi(scan.Text())
	for ; t > 0; t-- {
		scan.Scan()
		s := scan.Text()
		anagram := map[string]int{}
		for i := 1; i <= len(s)-1; i++ {
			stats := make([]int, 26)
			for _, r := range s[0:i] {
				idx := r - 'a'
				stats[idx]++
			}
			signature := fmt.Sprintf("%v", stats)
			anagram[signature]++
			for j := 1; j+i <= len(s); j++ {
				leaving := s[j-1] - 'a'
				coming := s[j+i-1] - 'a'
				stats[leaving]--
				stats[coming]++
				signature := fmt.Sprintf("%v", stats)
				anagram[signature]++
			}
		}
		cnt := 0
		for _, v := range anagram {
			cnt += v * (v - 1) / 2
		}
		fmt.Println(cnt)
	}
}
