package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
)

func parse(r io.Reader) ([]int, []int, error) {
	var left, right []int

	scan := bufio.NewScanner(r)
	for scan.Scan() {
		var l, r int
		_, err := fmt.Sscanf(scan.Text(), "%d %d", &l, &r)
		if err != nil {
			return nil, nil, err
		}

		left = append(left, l)
		right = append(right, r)
	}
	return left, right, nil
}

func part1(left, right []int) {
	slices.Sort(left)
	slices.Sort(right)

	var sum int
	for i := 0; i < len(left); i++ {
		diff := left[i] - right[i]
		if diff < 0 {
			sum += -diff
		} else {
			sum += diff
		}
	}

	fmt.Printf("Total distance: %d\n", sum)
}

func part2(left, right []int) {
	counts := map[int]int{}

	for _, val := range right {
		counts[val]++
	}

	var sum int
	for _, val := range left {
		sum += counts[val] * val
	}

	fmt.Println("Total similarity:", sum)
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	left, right, err := parse(bufio.NewReader(f))
	if err != nil {
		panic(err)
	}

	part1(left, right)

	part2(left, right)

	return
}
