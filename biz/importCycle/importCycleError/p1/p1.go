package p1

import (
	"fmt"
	"others/go-demos/biz/importCycle/importCycleError/p2"
)

type P1 struct {
}

func NewP1() *P1 {
	return &P1{}
}

func (p *P1) PrintP1() {
	fmt.Println("Hello, This is P1")
}

func (p *P1) DoSomething() {
	p22 := p2.NewP2()
	p22.PrintP2()
}
