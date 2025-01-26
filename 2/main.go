package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// в лоб
func countDigits(n int) int {
	if n < 10 {
		return n + 1
	}

	if n >= 10 && n < 100 {
		return 10 + n/11
	}

	if n >= 100 && n < 1000 {
		return 20 + n/111
	}

	if n >= 1000 && n < 10000 {
		return 30 + n/1111
	}

	if n >= 10000 && n < 100000 {
		return 40 + n/11111
	}

	if n >= 100000 && n < 1000000 {
		return 50 + n/111111
	}

	if n >= 1000000 && n < 10000000 {
		return 60 + n/1111111
	}

	if n >= 10000000 && n < 100000000 {
		return 70 + n/11111111
	}

	if n >= 100000000 && n < 1000000000 {
		return 80 + n/111111111
	}

	return 90 + n/1111111111
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())

	pineapples := make([]int, t)

	for i := 0; i < t; i++ {
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		pineapples[i] = n
	}

	for _, n := range pineapples {
		result := countDigits(n)
		fmt.Println(result)
	}
}
