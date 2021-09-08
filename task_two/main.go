package main

import (
	"fmt"
"io/ioutil"
"strconv"
"strings"
)


/*

https://contest.yandex.ru/contest/19036/problems/C/

 */

func main() {
	var k, n int
	dat, _ := ioutil.ReadFile("./input.txt")
	lstring := string(dat)
	nuf := strings.Split(lstring, "\n")
	params := strings.Split(nuf[0], " ")
	k, _ = strconv.Atoi(params[0])
	n, _ = strconv.Atoi(params[1])

	numsbers := strings.Split(nuf[1], " ")
	pPoints, vPoints := 0, 0
	for idx := 0; idx < n; idx++ {
		digit, _ := strconv.Atoi(numsbers[idx])
		if digit%15 == 0 {
			continue
		}
		delThree := digit % 3
		delFive := digit % 5
		if delThree == 0 && delFive == 0 {
			continue
		}
		if delThree == 0 {
			pPoints++
		} else if delFive == 0 {
			vPoints++

		}
		if vPoints == k || pPoints == k {
			break
		}
	}
	if vPoints > pPoints {
		fmt.Println("Vasya")
	} else if pPoints > vPoints {
		fmt.Println("Petya")
	} else {
		fmt.Println("Draw")
	}
}



