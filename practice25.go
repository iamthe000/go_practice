package main

import (
	"runtime/debug"
)

func main() {
	functionA()
}

func functionA() {
	functionB()
}

func functionB() {
	debug.PrintStack()
}
