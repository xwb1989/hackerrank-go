package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scan := func() func() int {
		scan := bufio.NewScanner(os.Stdin)
		scan.Split(bufio.ScanWords)
		return func() int {
			scan.Scan()
			i, _ := strconv.Atoi(scan.Text())
			return i
		}
	}()

	t := scan()

	for ; t > 0; t-- {
		ns := scan()
		es := scan()
		g := make([][]int, ns)
		for ; es > 0; es-- {
			a := scan() - 1
			b := scan() - 1
			g[a] = append(g[a], b)
			g[b] = append(g[b], a)
		}

		start := scan() - 1

		visited := make([]bool, ns)
		dist := make([]int, ns)
		for i, _ := range dist {
			dist[i] = -1
		}
		queue := list.New()
		queue.PushBack(start)
		visited[start] = true
		dist[start] = 0
		for queue.Len() != 0 {
			curr := queue.Remove(queue.Front()).(int)
			next := g[curr]
			for _, n := range next {
				if !visited[n] {
					visited[n] = true
					queue.PushBack(n)
					dist[n] = dist[curr] + 1
				}
			}
		}

		sep := ""
		for _, d := range dist {
			if d != 0 {
				fmt.Print(sep)
				if d != -1 {
					fmt.Print(d * 6)
				} else {
					fmt.Print(-1)
				}
				sep = " "
			}
		}
		fmt.Println()
	}
}
