package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func main() {
	task12()
}

func sort_r(numbers []int, wg *sync.WaitGroup) {
	sort.Ints(numbers)
	fmt.Println("Sorted part:", numbers)
	wg.Done()
}

func merge(a, b []int) []int {
	result := make([]int, 0, len(a)+len(b))

	i, j := 0, 0

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}

	result = append(result, a[i:]...)
	result = append(result, b[j:]...)

	return result
}

func task12() {
	fmt.Println("Input a series of integer")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	parts := strings.Split(s, " ")
	var data []int
	var wg sync.WaitGroup

	for _, v := range parts {
		d, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("Write a correct data")
			return
		}
		data = append(data, d)
	}

	n := len(data)
	chunkSize := n / 4

	part1 := data[0:chunkSize]
	part2 := data[chunkSize : 2*chunkSize]
	part3 := data[2*chunkSize : 3*chunkSize]
	part4 := data[3*chunkSize:]

	wg.Add(4)

	go sort_r(part1, &wg)
	go sort_r(part2, &wg)
	go sort_r(part3, &wg)
	go sort_r(part4, &wg)

	wg.Wait()

	firstHalf := merge(part1, part2)
	secondHalf := merge(part3, part4)

	sortedData := merge(firstHalf, secondHalf)

	fmt.Println(sortedData)

}
