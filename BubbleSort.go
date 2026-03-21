package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	task7()
}

func BubbleSort(numbers []int) {

	for i := 0; i < len(numbers)-1; i++ {
		for j := 0; j < len(numbers)-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				Swap(numbers, j)
			}
		}
	}
}
func Swap(numbers []int, index int) {
	temp := numbers[index]
	numbers[index] = numbers[index+1]
	numbers[index+1] = temp
}
func task7() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Type up to 10 integers")
	scanner.Scan()
	s1 := scanner.Text()
	numbersStr := strings.Split(s1, " ")
	numbers := make([]int, 0, 10)
	for _, val := range numbersStr {
		temp, err := strconv.Atoi(val)
		if err != nil {
			break
		}
		numbers = append(numbers, temp)
	}
	BubbleSort(numbers)
	fmt.Println(numbers)
}
