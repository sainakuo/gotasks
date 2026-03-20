package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func task6() {
	type Name struct {
		fname string
		lname string
	}

	sli := make([]Name, 0)

	var fileName string

	fmt.Println("Enter the file name")

	scanner1 := bufio.NewScanner(os.Stdin)

	scanner1.Scan()

	fileName = scanner1.Text()

	f, _ := os.Open(fileName)

	scanner2 := bufio.NewScanner(f)

	for scanner2.Scan() {
		line := scanner2.Text()
		parts := strings.Split(line, " ")

		p1 := Name{fname: parts[0], lname: parts[1]}

		sli = append(sli, p1)
	}

	for _, v := range sli {
		fmt.Println(v.fname, v.lname)
	}
}
