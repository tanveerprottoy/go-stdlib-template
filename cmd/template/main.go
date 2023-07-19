package main

import "github.com/tanveerprottoy/stdlib-go-template/internal/app/template"

func main() {
	a := template.NewApp()
	a.Run()
}

func Multiply(x, y int) int {
	return x * y
}
