package main

import "test/testpackage"

func main() {
	testpackage.MyFunction("1")
	testpackage.MyFunction("This is a string  ")
	testpackage.MyFunction("Unsupported step")
}
