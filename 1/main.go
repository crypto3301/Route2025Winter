package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type SizeRoom struct {
	m int
	n int
}

type LightAndLine struct {
	x   int
	y   int
	dir string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())

	data := make([]SizeRoom, t)

	for i := 0; i < t; i++ {
		scanner.Scan()
		item := scanner.Text()
		parts := strings.Fields(item)
		m, _ := strconv.Atoi(parts[0])
		n, _ := strconv.Atoi(parts[1])
		data[i] = SizeRoom{m: m, n: n}
	}

	for _, room := range data {
		var lightandline []LightAndLine

		if room.m == 1 && room.n == 1 {
			lightandline = append(lightandline, LightAndLine{x: 1, y: 1, dir: "R"})
		} else if room.m == 1 {
			lightandline = append(lightandline, LightAndLine{x: 1, y: 1, dir: "R"})
		} else if room.n == 1 {
			lightandline = append(lightandline, LightAndLine{x: 1, y: 1, dir: "D"})
		} else if room.m > room.n {
			lightandline = append(lightandline, LightAndLine{x: 1, y: 1, dir: "D"})
			lightandline = append(lightandline, LightAndLine{x: room.m, y: room.n, dir: "U"})
		} else if room.n > room.m {
			lightandline = append(lightandline, LightAndLine{x: 1, y: 1, dir: "R"})
			lightandline = append(lightandline, LightAndLine{x: room.m, y: room.n, dir: "L"})
		} else {
			lightandline = append(lightandline, LightAndLine{x: 1, y: 1, dir: "R"})
			lightandline = append(lightandline, LightAndLine{x: room.m, y: room.n, dir: "L"})
		}

		fmt.Println(len(lightandline))

		for _, f := range lightandline {
			fmt.Printf("%d %d %s\n", f.x, f.y, f.dir)
		}
	}
}

/*идея моя такая: находим самую длинную сторону и ставим вдоль нее 1 фонарик,
тогда в противоположный угол комнат ставим другой фонарик в сторону первого фонарика*/
