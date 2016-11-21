package main

import (
	"fmt"
	"strings"
)

func main() {
	module := "function basics"
	author := "ray masson"

	fmt.Println(converter(module, author))

	bestFinish := bestLeagueFinishes(13, 10, 13, 17, 14, 16, 7, 5, 2)

	fmt.Println(bestFinish)
}

func converter(s1, s2 string) (str1, str2 string) {
	s1 = strings.Title(s1)
	s2 = strings.ToUpper(s2)

	return s1, s2
}

func bestLeagueFinishes(finishes ...int) int {
	best := finishes[0]

	for _, i := range finishes {
		if i < best {
			best = i
		}
	}

	return best
}
