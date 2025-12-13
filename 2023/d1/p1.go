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
	var res int
	for _, li := range in {
		fmt.Println(li)
		var n1 string
		var n2 string

		for i := 0; i < len(li); i++ {
			_, err := strconv.Atoi(string(li[i]))
			if err != nil { continue }
			n1 = string(li[i])
			break
		}
		fmt.Printf("  %s\n", n1)
		
		for i := len(li)-1; i > -1; i-- {
			_, err := strconv.Atoi(string(li[i]))
			if err != nil { continue }
			n2 = string(li[i])
			break
		}
		fmt.Printf("  %s\n", n2)

		f, _ := strconv.Atoi(n1+n2)
		res += f
	}

	fmt.Println(res)
}
