package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Direction int

const (
	DOWN Direction = iota
	RIGHT
	UP
	LEFT
)

func (d Direction) next() Direction {
	switch d {
	case DOWN:
		return RIGHT
	case RIGHT:
		return UP
	case UP:
		return LEFT
	case LEFT:
		return DOWN
	}
	return d
}

func locator(x, y, up, left, down, right, r int, d Direction) (int, int) {
	if r == 0 {
		return x, y
	} else {
		switch d {
		case DOWN:
			if x+r <= down {
				return x + r, y
			} else {
				return locator(down, y, up, left, down, right, r-(down-x), RIGHT)
			}
		case RIGHT:
			if y+r <= right {
				return x, y + r
			} else {
				return locator(x, right, up, left, down, right, r-(right-y), UP)
			}
		case UP:
			if x-r >= up {
				return x - r, y
			} else {
				return locator(up, y, up, left, down, right, r-(x-up), LEFT)
			}
		case LEFT:
			if y-r >= left {
				return x, y - r
			} else {
				return locator(x, left, up, left, down, right, r-(y-left), DOWN)
			}
		}
		return x, y
	}
}

func main() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Split(bufio.ScanLines)

	scan.Scan()
	parts := strings.Split(scan.Text(), " ")

	atoi := func(s string) int {
		i, _ := strconv.Atoi(s)
		return i
	}
	m, n, r := atoi(parts[0]), atoi(parts[1]), atoi(parts[2])

	matrix := make([][]string, m)
	for i := 0; i < m; i++ {
		scan.Scan()
		row := strings.Split(scan.Text(), " ")
		matrix[i] = row
	}

	res := make([][]string, m)
	for i := range res {
		res[i] = make([]string, n)
	}

	for i := range res {
		for j := range res[i] {
			var xd, yd int
			if i < m-1-i {
				xd = i
			} else {
				xd = m - 1 - i
			}
			if j < n-1-j {
				yd = j
			} else {
				yd = n - 1 - j
			}
			var ring int
			if xd < yd {
				ring = xd
			} else {
				ring = yd
			}
			var up, left, down, right int
			var d Direction
			up = ring
			down = m - 1 - ring
			left = ring
			right = n - 1 - ring
			rotate := r % ((down-up)*2 + (right-left)*2)
			//fmt.Printf("i:%v, j:%v, ring:%v, up:%v, left:%v, down:%v, right:%v, r:%v\n", i, j, ring, up, left, down, right, rotate)
			if i == up {
				if j == left {
					d = DOWN
				} else {
					d = LEFT
				}
			} else if i == down {
				if j == right {
					d = UP
				} else {
					d = RIGHT
				}
			} else {
				if j == left {
					d = DOWN
				} else {
					d = UP
				}
			}
			x, y := locator(i, j, up, left, down, right, rotate, d)
			res[x][y] = matrix[i][j]
		}
	}
	for _, row := range res {
		for _, cell := range row {
			fmt.Print(cell, " ")
		}
		fmt.Println()
	}
}
