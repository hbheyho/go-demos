package main

import (
	"fmt"
	"others/go-demos/biz/importCycle/importCycleError/p1"
	"others/go-demos/biz/importCycle/importCycleError/p2"
)

func main() {
	fmt.Println("Import Cycle Testing....")
	p1.NewP1().DoSomething()
	p2.NewP2().DoSomething()
}
