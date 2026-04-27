package main

import (
	"fmt"
	"strings"
)

func rotateString(s string, goal string) bool {
    return len(s) == len(goal) && strings.Contains(s + s, goal)
}

func main() {
	fmt.Println(rotateString("adc", "cad"))
}