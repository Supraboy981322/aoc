package main

import (
	"os"
	"fmt"
	"strconv"
	"strings"
)

var (
	maxC = map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
)

func main() {
	inR, err := os.ReadFile("p1.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read file: %v\n", err)
		os.Exit(1)
	}
	games := strings.Split(string(inR), "\n")
	games = games[:len(games)-1]
	var res int
	for i, gameR := range games {
		game := strings.Split(gameR, ":")[1]
		sets := strings.Split(game, ";")
		valid := true
		for _, set := range sets {
			for _, cubeR := range strings.Split(set, ",") {
				cube := cubeR[1:]
				col := strings.Split(cube, " ")[1]
				val, _ := strconv.Atoi(strings.Split(cube, " ")[0])
				if maxC[col] < val {
					valid = false
				}
			}
		}
		if valid { res += i+1 }
	}
	fmt.Printf("%d\n", res)
}
