package main

import (
	"fmt"
	"others/go-demos/biz/importCycle/importCycleBroken/p1"
	"others/go-demos/biz/importCycle/importCycleBroken/p2"
)

func main() {
	fmt.Println("Import Cycle testing...")
	p11 := p1.NewP1()
	p11.DoSomething()
	p2.NewP2(p11).DoSomething()
}
