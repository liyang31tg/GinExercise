package main

import "fmt"

type T int

func IsClosed(ch <-chan T) bool {
	select {
	case <-ch:
		fmt.Println("会执行到这里吗")
		return true
	default:
	}

	return false
}

func main() {
	c := make(chan T)
	fmt.Println(IsClosed(c)) // false
	close(c)
	v, ok := <-c
	if !ok {
		fmt.Println(ok, v)
	}
	fmt.Println(IsClosed(c)) // true
}
