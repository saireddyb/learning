package main

import (
	"fmt"
	"runtime"
)

func main(){
	version := runtime.Version()
	fmt.Print(version)
}