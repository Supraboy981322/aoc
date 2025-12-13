/*Getting wrong result, not sure why,
 * when manually checking the values on
 *  each line, they seem correct, so I'm
 *   probably just miss-interpretting instructions
 *    (which is very likely)
 */

package main

import (
	"os"
	"fmt"
	"strconv"
	"strings"
)

var (
	nums = map[string]any{
		"o": map[string]any{
			"n": map[string]any{"e": 1,},
		},
		"t": map[string]any{
			"w": map[string]any{"o": 2,},
			"h": map[string]any{"r": map[string]any{"e":map[string]any{"e": 3,},},},
		},
		"f": map[string]any{
			"o": map[string]any{"u": map[string]any{"r": 4,},},
      "i": map[string]any{"v": map[string]any{"e": 5,},},
		},
		"s": map[string]any{
      "i": map[string]any{"x":6,},
			"e": map[string]any{"v": map[string]any{"e": map[string]any{"n": 7,},},},
		},
		"e": map[string]any{
			"i": map[string]any{"g": map[string]any{"h": map[string]any{"t": 8,},},},
		},
		"n": map[string]any{
			"i": map[string]any{"n": map[string]any{"e": 9,},},
		},
	}
)

func main() {
	//read file
	inR, _ := os.ReadFile("in.txt")

	//split into string slice of lines
	in := strings.Split(string(inR), "\n")

	//for debugging (changed to, for example, in = in[5:6])
	in = in

	var res int //result var

	//iterate through each line
	for k, l := range in {
		//if the line is empty, skip
		if l == "" { continue }
		
		//print the line
		fmt.Println("\n\033[34m"+l+"\033[0m")
		var n1 string //first number
		var n2 string //final number

		//parse line twice
		for i := 0; i < 2; i++ {
			var cur string //current character

			m := nums //start-out with default map

			//iterate through each char in line
			for j := 0; j < len(l); j++ {
				cur = string(l[j]) //set the current char
				//check if cur is end of map
				if _, ok := m[cur].(int); !ok {
					//if not, check if it's the next key map
					mT, ok := m[cur].(map[string]any)
					if ok { //if it is the next key 
						fmt.Print("\033[32m"+cur) //print it in green
						m = mT //set map to map at current char
						continue //skip rest of iteration
					} else { //if it's not the next key
						fmt.Print("\033[31m"+cur) //print it in red
						m = nums //reset map
						if j-1 > -1 { //check if previous char is in root of map
							if mT, ok := m[string(l[j-1])].(map[string]any); ok {
								//if so,
								//  check if current char is
								//    the next val of that map
								if mT, ok := mT[cur].(map[string]any); ok {
									m = mT //move the map to that position if so
									fmt.Print("\b\033[32m"+cur) //change previous char to green
									continue //skip rest of iteration
								}
							}
						}
					}
				} else { fmt.Print("\033[32m"+cur) } //print cur in green

				//double check that it's an int
				//  (it sometimes isn't at this point)
				nR, ok := m[cur].(int)
				m = nums //reset the map
				if ok { //if it is an int
					//convert it to a string
					n := strconv.Itoa(nR)
//					j--  //move current char back by one
					//if iteration 1 of line, set num 1
					//  otherwise, set num 2
					if i == 0 { n1 = n; break //stop loop 
  				} else { n2 = n }
					continue //skip rest of iteration
				}

				//check if it can be converted to int
				_, err := strconv.Atoi(cur)
				//skip rest of iteration if not
				if err != nil { continue }
				
				//if it can be an converted to int,
				//  print it in yellow
				fmt.Print("\b\033[33m"+cur)

				//make sure the num for current
				//  iteration is set
				if i == 0 { n1 = cur; break
				} else { n2 = cur }
			}
			//print end of iteration
			fmt.Printf("\n\033[0miteration %d of line %d done\n", i+1, k+1)
		}

		//convert concatted string nums to an int
		f, _ := strconv.Atoi(n1+n2)

		//print vals for debugging 
		fmt.Printf("  n1: %s\n  n2: %s\n  f: %d\n", n1, n2, f)

		//add final int num to result
		res += f
	}

	//print result
	fmt.Printf("\nres:  %d\n", res)
}
