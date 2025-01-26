package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	testCases := make([]struct {
		n     int
		words []string
	}, t)

	for testCase := 0; testCase < t; testCase++ {
		var n int
		fmt.Fscan(in, &n)

		words := make([]string, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &words[i])
		}

		testCases[testCase] = struct {
			n     int
			words []string
		}{n, words}
	}

	for _, testCase := range testCases {
		words := testCase.words

		fullEquivalentWords := make(map[string]int)
		unevenEquivalentWords := make(map[string]int)
		evenEquivalentWords := make(map[string]int)

		for _, word := range words {
			fullEquivalentWords[word]++

			var unevenBuilder, evenBuilder strings.Builder
			for i := 0; i < len(word); i++ {
				if i%2 == 1 {
					unevenBuilder.WriteByte(word[i])
				} else {
					evenBuilder.WriteByte(word[i])
				}
			}
			unevenWord := unevenBuilder.String()
			evenWord := evenBuilder.String()

			if len(word) == 1 {
				evenEquivalentWords[evenWord]++
			} else {
				unevenEquivalentWords[unevenWord]++
				evenEquivalentWords[evenWord]++
			}
		}

		sum := 0
		for _, count := range unevenEquivalentWords {
			sum += count * (count - 1) / 2
		}
		for _, count := range evenEquivalentWords {
			sum += count * (count - 1) / 2
		}

		for word, count := range fullEquivalentWords {
			if len(word) > 1 {
				sum -= count * (count - 1) / 2
			}
		}

		fmt.Fprintln(out, sum)
	}
}
