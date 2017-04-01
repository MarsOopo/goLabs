package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("vim-go")
	fmt.Println(strconv.Itoa(100))
	i, _ := strconv.Atoi("aaa")
	fmt.Println(i)
}
