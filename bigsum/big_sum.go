package bigsum

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Split(bufio.ScanWords)
	scan.Scan()
	cnt, _ := strconv.Atoi(scan.Text())
	var sum int64
	for i := 0; i < cnt; i++ {
		scan.Scan()
		j, _ := strconv.ParseInt(scan.Text(), 10, 64)
		sum += j
	}
	fmt.Println(sum)
}
