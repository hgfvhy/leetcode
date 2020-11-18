package main

import (
	"fmt"
	"math"
)

func main() {
	i := 1 << 32
	temp := int32(i)
	fmt.Println(reverse(i-1), math.MaxInt32, temp)
}

func reverse(x int) int {
	var res int
	for x != 0 {
		if temp := int32(res); (temp*10)/10 != temp {
			return 0
		}
		res = res*10 + x%10
		x = x / 10
	}
	return res
}
