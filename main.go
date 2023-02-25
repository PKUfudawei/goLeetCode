package main

import (
	"fmt"
	goLeetCode "goLeetCode/easy"
	"reflect"
)

func main() {
	var result = []int{}
	result = nil
	fmt.Println(reflect.TypeOf(result), reflect.ValueOf(result))
	result1 := append([]int{1, 2, 3}, result...)
	fmt.Println(reflect.TypeOf(result1), reflect.ValueOf(result1))
	fmt.Println(goLeetCode.ClimbStairs(5))
}
