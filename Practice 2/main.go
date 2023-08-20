package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func spread(num ...string) {
	fmt.Println(num)
}

func matrixMaker(col, row int) [][]int {
	matrix := make([][]int, 0)

	for i := 0; i <= col; i++ {
		row_item := make([]int, 0)
		for j := 0; j <= row; j++ {
			row_item = append(row_item, i*j)
		}
		matrix = append(matrix, row_item)
	}
	return matrix
}

func getOnlineUsers(users []string) map[string]map[string]int {
	onlineUsers := make(map[string]map[string]int)

	for _, user := range users {
		if user == "" {
			continue

		}
		firstChar := user[0]
		_, ok := onlineUsers[string(firstChar)]
		if !ok {
			onlineUsers[string(firstChar)] = make(map[string]int)
		}
		onlineUsers[string(firstChar)][user]++
	}
	return onlineUsers
}

func main() {
	cores := runtime.NumCPU()
	fmt.Println(cores)
	users := []string{"alvan", "mike", "alex"}
	fmt.Println(getOnlineUsers(users))
	fmt.Println("=============**Matrix**=============")
	fmt.Println(matrixMaker(5, 5))
	fmt.Println("=============**spread operators**=============")
	nums := []string{"1", "2", "3"}
	spread(nums...)

	fmt.Println("=============**interfaces**=============")
	s := rect{2, 4}
	fmt.Println(s.area())

	checkType(s)

	values := reflect.ValueOf(s)
	types := values.Type()
	for i := 0; i < values.NumField(); i++ {
		fmt.Println(types.Field(i).Index[0], types.Field(i).Name, values.Field(i))
	}

}
