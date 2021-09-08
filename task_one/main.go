package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

/*
Описание задачи
https://contest.yandex.ru/contest/19036/problems/B/

 */


const MaxUint = ^uint64(0)
const MaxInt = int(MaxUint >> 1)

func main() {
	var n, x, k int

	fmt.Scanf("%d %d %d", &n, &x, &k)
	reader := bufio.NewReader(os.Stdin)
	numbers:= read(reader,n)
	filteredNumbers := make([]int, 0)
	tempmap := make(map[int][]int)
	for _, val := range numbers {
		delimiter := val % x
		_, ok := tempmap[delimiter]
		if ok == false {
			tempmap[delimiter] = make([]int, 0)
			tempmap[delimiter] = append(tempmap[delimiter], val)
		} else {
			tempmap[delimiter] = append(tempmap[delimiter], val)
		}
	}
	for _, val := range tempmap {
		sort.Ints(val)
		filteredNumbers = append(filteredNumbers, val[0])
	}
	sort.Ints(filteredNumbers)
	fmt.Println(binarySearch(x, filteredNumbers, k))

}



func binarySearch(repeat int, numbs []int, k int) int {
	minValue := 0
	maxValue := MaxInt
	for {
		if minValue < maxValue {
			middle := (minValue + maxValue) / 2
			count := getAllRingersCount(middle, repeat, numbs)
			if count < k {
				minValue = middle + 1
			} else {
				maxValue = middle
			}
		} else {
			return minValue
		}
	}
}

func getAllRingersCount(time, repeat int, numbs []int) int {
	count := 0
	for _, val := range numbs {
		count += findRingedCount(time, val, repeat)
		if count == 0 {
			return 0
		}
	}
	return count
}

func findRingedCount(time, ringerFirst, repeat int) int {
	if time < ringerFirst {
		return 0
	}
	count := (time - ringerFirst) / repeat
	count++
	return count
}


func read (reader *bufio.Reader, n int)([]int) {

	a := make([]int, n)
	for i:=0; i<n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	return a
}