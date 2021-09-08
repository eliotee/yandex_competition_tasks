package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)


/*

https://contest.yandex.ru/contest/19036/problems/D/

 */
const maxRequests = 25

func main() {
	var commitsCount int

	fmt.Scan(&commitsCount)

	middle := commitsCount / 2
	maxCommit := commitsCount
	minCommit := 1
	cachedReq := 0
	requests := 0
	for {
		requests++
		if (requests == maxRequests){
			fmt.Println("Fail!")
			return
		}
		if (commitsCount == 1){
			printWithFlush("! 1")
		}
		answer := communicate(middle)
		if answer == 1 {
			minCommit = middle
			cachedReq = middle
			middle = (maxCommit + middle) / 2
		} else if answer == 0 {
			maxCommit = middle
			middle = (minCommit + middle) / 2
		}
		//4751
		if math.Abs(float64(minCommit-maxCommit)) == 1 {
			cachedReq++
			printWithFlush("! " + strconv.Itoa(cachedReq))
			break
		}

	}
}

func communicate(number int) int {
	var answer int
	printWithFlush(strconv.Itoa(number))
	fmt.Scanln(&answer)


	return answer
}


func printWithFlush(s string){
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	f.Write([]byte(s+ "\n"))
}