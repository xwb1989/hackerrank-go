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

	h := next()
	m := next()

	switch {
	case m == 0:
		fmt.Printf("%s o' clock\n", toWords(h))
	case m == 15:
		fmt.Printf("quarter past %s\n", toWords(h))
	case m < 30:
		s := ""
		if m > 1 {
			s = "s"
		}
		fmt.Printf("%s minute%s past %s\n", toWords(m), s, toWords(h))
	case m == 30:
		fmt.Printf("half past %s\n", toWords(h))
	case m == 45:
		fmt.Printf("quarter to %s\n", toWords(h+1))
	default:
		s := ""
		if 60-m > 1 {
			s = "s"
		}
		fmt.Printf("%s minute%s to %s\n", toWords(60-m), s, toWords(h+1))
	}
}

func toWords(m int) string {
	dict := map[int]string{
		1:  "one",
		2:  "two",
		3:  "three",
		4:  "four",
		5:  "five",
		6:  "six",
		7:  "seven",
		8:  "eight",
		9:  "nine",
		10: "ten",
		11: "eleven",
		12: "twelve",
		13: "thirteen",
		14: "fourteen",
		15: "fifteen",
		16: "sixteen",
		17: "seventeen",
		18: "eighteen",
		19: "nineteen",
		20: "twenty",
		30: "thirty",
		40: "forty",
		50: "fifty",
	}
	if m <= 20 {
		return dict[m]
	} else {
		return fmt.Sprintf("%s %s", dict[m/10*10], dict[m%10])
	}
}
