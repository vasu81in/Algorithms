/*
*
* Vasu Mahalingam, 2017
*
* Sum of all  sub arrays
* -----------------------
* Input   : arr[] = {1, 2, 3}
* Output  : 20
* Explanation : {1} + {2} + {3} + {2 + 3} + 
*               {1 + 2} + {1 + 2 + 3} = 20
*
* Input  : arr[] = {1, 2, 3, 4}
* Output : 50
*
*/


package main

import (
	"fmt"
)

func subset(a []int) {
	var i, x, s, result int
	for i, _ = range a {
		s = 0
		for x = i; x < len(a); {
			fmt.Println("( ", a[x], " )")
			s += a[x]
			fmt.Println("Sum :", s)
			result += s
			x++
		}
	}
	fmt.Println("Result", result)
}

func main() {
	subset([]int{1, 2, 3, 4})
	subset([]int{1, 2, 3})
}

