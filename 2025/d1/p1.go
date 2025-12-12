package main

import (
	"os"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	inR, _ := os.ReadFile("in.txt")
	in := strings.Split(string(inR), "\n")
	in = in[:len(in)-1]
	var res int
	dir := 50
	_ = dir
	for i, l := range in {
		_ = i 
		d := l[0]; _ = d
		nR, _ := strconv.Atoi(l[1:])
		n := nR 
		switch d {
     case 'L':
			for j := 0; j < n; j++ {
				dir--
				if dir == -1 { dir = 99 }
			}
		 case 'R': 
			for j := 0; j < n; j++ {
				dir++
				if dir == 100 { dir = 0 }
			}
     default:
			fmt.Printf("bad dir: %s (line %d)", string(d), i)
			os.Exit(1)
		}
		if dir == 0 { res++ }
	}
	fmt.Printf("\nres:  %d\n",res)
}
