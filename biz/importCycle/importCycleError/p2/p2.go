package p2

import (
	"fmt"
	"others/go-demos/biz/importCycle/importCycleError/p1"
)

type P2 struct {
}

func NewP2() *P2 {
	return &P2{}
}

func (p *P2) PrintP2() {
	fmt.Println("Hello, This is P2")
}

func (p *P2) DoSomething() {
	p11 := p1.NewP1()
	p11.PrintP1()
}
