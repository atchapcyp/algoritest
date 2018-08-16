package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//domino()
	// maps := [6][6]string{{"0", "0", "*", "0", "0", "0"},
	// 	{"0", "*", "0", "0", "*", "0"},
	// 	{"0", "0", "0", "0", "*", "0"},
	// 	{"*", "*", "*", "0", "*", "0"},
	// 	{"*", "*", "0", "0", "*", "0"},
	// 	{"*", "*", "0", "*", "0", "0"}}
	var route string
	mapss := [6][6]rune{{'0', '0', '*', '0', '0', '0'},
		{'0', '*', '0', '0', '*', '0'},
		{'0', '0', '0', '0', '*', '0'},
		{'*', '*', '*', '0', '*', '0'},
		{'*', '*', '0', '0', '*', '*'},
		{'*', '*', '0', '*', '0', '0'}}
	showmap := maze(mapss)
	printShowMap(showmap)
	move(mapss, showmap, 0, 0, route)

	if mapss[1][0] == '*' {

		println(mapss[1][0])
	} else {
		println("Not found")
	}
}

func ezatoi(a string, b, c int) (int, error) {
	return strconv.Atoi(a[b:c])
}

func domino() {
	s := "409010020000100"
	startint := 0
	pushcount := 0
	for i := 0; i < len(s); i++ {
		if startint > 0 {
			startint--
		}
		startvalue, _ := ezatoi(s, i, i+1)
		if startint == 0 && startvalue > 0 {
			startint = startvalue
			pushcount++
		}

		if startvalue >= startint && startint != 0 { // keep fall
			startint = startvalue
			startint++
		}
	}
	fmt.Printf("You need to push %d time(s) ", pushcount)
}

func factorial(i int) int {
	if i <= 1 {
		return 1
	}
	return i * factorial(i-1)
}

func fibonuci(i int) int {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fibonuci(i-1) + fibonuci(i-2)
}

func maze(mapss [6][6]rune) [6][6]string {
	char := 'A'
	showmap := [6][6]string{}
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			if mapss[i][j] == '0' {
				showmap[i][j] = string(char)
				char++
			} else {
				showmap[i][j] = "*"
			}
		}
	}
	return showmap
}

func printShowMap(showmap [6][6]string) {
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			fmt.Printf("%v", showmap[i][j])
			if j == 5 {
				println()
			}
		}
	}
}

func move(walk [6][6]rune, s [6][6]string, x, y int, r string) bool {
	r += s[x][y]
	println("Add route : " + r)
	if x == 5 && y == 5 {
		return true
	} else {
		if x+1 < 6 && walk[x+1][y] == '0' && !strings.Contains(r, s[x+1][y]) {
			fmt.Printf(" x = %d , y = %d , walk[%d][%d] \n", x, y, x+1, y)
			if move(walk, s, x+1, y, r) {
				return true
			}
		}
		if y+1 < 6 && walk[x][y+1] == '0' && !strings.Contains(r, s[x][y+1]) {
			fmt.Printf(" x = %d , y = %d , walk[%d][%d] \n", x, y, x, y+1)
			if move(walk, s, x, y+1, r) {
				return true
			}

		}

		if x-1 >= 0 && walk[x-1][y] == '0' && !strings.Contains(r, s[x-1][y]) {
			fmt.Printf(" x = %d , y = %d , walk[%d][%d] \n", x, y, x-1, y)
			if move(walk, s, x-1, y, r) {
				return true
			}
		}

		if y-1 >= 0 && walk[x][y-1] == '0' && !strings.Contains(r, s[x][y-1]) {
			fmt.Printf(" x = %d , y = %d , walk[%d][%d] \n", x, y, x, y-1)
			if move(walk, s, x, y-1, r) {
				return true
			}
		}
		r = r[0 : len(r)-1]
		walk[x][y] = '*'
		println("Remove route : " + r)
		return false
	}
}

type coordinate struct {
	x, y  int
	value string
}
