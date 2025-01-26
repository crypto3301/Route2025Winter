package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TestCase struct {
	NumProducts int
	Products    map[string]int
	PriceCount  map[int]int
	Query       string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	testCases := make([]TestCase, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		num, _ := strconv.Atoi(scanner.Text())

		products := make(map[string]int)
		priceCount := make(map[int]int)

		for j := 0; j < num; j++ {
			scanner.Scan()
			parts := strings.Fields(scanner.Text())
			name := parts[0]
			price, _ := strconv.Atoi(parts[1])
			products[name] = price
			priceCount[price]++
		}

		scanner.Scan()
		query := scanner.Text()

		testCases[i] = TestCase{
			NumProducts: num,
			Products:    products,
			PriceCount:  priceCount,
			Query:       query,
		}
	}

	for _, testCase := range testCases {
		pairs := strings.Split(testCase.Query, ",")

		usedPrice := make(map[int]bool)    //cost used
		usedNames := make(map[string]bool) //name used

		valid := true

		for _, pair := range pairs {
			parts := strings.Split(pair, ":")
			if len(parts) != 2 {
				valid = false
				break
			}
			name := parts[0]
			priceStr := parts[1]

			if len(priceStr) > 1 && priceStr[0] == '0' {
				valid = false
				break
			}

			price, err := strconv.Atoi(priceStr)
			if err != nil {
				valid = false
				break
			}

			if expectedPrice, exists := testCase.Products[name]; !exists || expectedPrice != price {
				valid = false
				break
			}

			if usedNames[name] {
				valid = false
				break
			}

			if usedPrice[price] {
				valid = false
				break
			}

			usedPrice[price] = true
			usedNames[name] = true
		}

		if valid {
			for price, count := range testCase.PriceCount {
				if count >= 1 && !usedPrice[price] {
					valid = false
					break
				}
			}
		}

		if valid {
			fmt.Println("YES")
		} else {

			fmt.Println("NO")
		}
	}
}
