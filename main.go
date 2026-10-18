package main

import (
	"fmt"
)

func main() {
	s := NewStore()
	s.Set("a", "42")
	s.Set("b", "72")

	value, _ := s.Get("a")

	fmt.Println(value)
	fmt.Println(s.data)
	s.Delete("b")
	fmt.Println(s.data)

	fmt.Println("Gokv - Go key value store project")
}
