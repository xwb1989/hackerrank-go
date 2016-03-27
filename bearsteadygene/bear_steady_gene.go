package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://www.hackerrank.com/challenges/bear-and-steady-gene

func hit(cnt *[4]int, r byte) {
	switch r {
	case 'A':
		(*cnt)[0]++
	case 'T':
		(*cnt)[1]++
	case 'C':
		(*cnt)[2]++
	case 'G':
		(*cnt)[3]++
	}
}

func unhit(cnt *[4]int, r byte) {
	switch r {
	case 'A':
		(*cnt)[0]--
	case 'T':
		(*cnt)[1]--
	case 'C':
		(*cnt)[2]--
	case 'G':
		(*cnt)[3]--
	}
}

func valid(req, now *[4]int) bool {
	for i := range *req {
		if (*req)[i] < 0 {
			continue
		} else {
			if (*req)[i] > (*now)[i] {
				return false
			}
		}
	}
	return true
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 500001)
	line, _ := reader.ReadString('\n')
	l, _ := strconv.Atoi(strings.TrimSpace(line))

	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)

	cnt := [4]int{}
	for i := range s {
		hit(&cnt, s[i])
	}

	allZero := true
	for i := range cnt {
		cnt[i] = cnt[i] - l/4
		allZero = allZero && cnt[i] == 0
	}
	if allZero {
		fmt.Println(0)
		return
	}

	var now [4]int
	var i, j, res int
	res = l
	for ; i < l; i++ {
		hit(&now, s[i])
		for ; valid(&cnt, &now) && j < i; j++ {
			if i-j+1 < res {
				res = i - j + 1
			}
			unhit(&now, s[j])
		}
	}
	fmt.Println(res)
}
