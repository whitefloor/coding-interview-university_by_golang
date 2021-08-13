package main

import (
	"log"
)

func main() {
	v := abs(-126)
	log.Println(v)
}

func abs(x int) int {
	con := x >> 63

	return (x ^ con) - con
}
