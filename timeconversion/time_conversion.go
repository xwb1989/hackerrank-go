package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	s := scan.Text()

	time := strings.Split(s, ":")
	hour, _ := strconv.Atoi(time[0])
	if strings.HasSuffix(time[2], "PM") {
		if hour != 12 {
			hour += 12
		}
	} else {
		if hour == 12 {
			hour = 0
		}
	}
	time[0] = fmt.Sprintf("%02d", hour)
	time[2] = time[2][:len(time[2])-2]

	fmt.Println(strings.Join(time, ":"))
}
