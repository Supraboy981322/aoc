package main

import (
	"os"
	"fmt"
	"strconv"
	"strings"
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
	for _, gameR := range games {
		game := strings.Split(gameR, ":")[1]
		sets := strings.Split(game, ";")
		minNum := map[string]int{
			"blue":  0,
			"red":   0,
			"green": 0,
		}
		for _, set := range sets {
			for _, cubeR := range strings.Split(set, ",") {
				cube := strings.Split(cubeR[1:], " ")
				col := cube[1]
				val, _ := strconv.Atoi(cube[0])
				if val > minNum[col] {
					minNum[col] = val
				}
			}
		}
		toRes := 1
		for _, n := range minNum {
			toRes = n * toRes
		}
		res += toRes
	}
	fmt.Printf("%d\n", res)
}
