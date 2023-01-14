package p1

import (
	"fmt"
	"others/go-demos/biz/importCycle/importCycleBroken/p2"
)

type P1 struct {
}

func NewP1() *P1 {
	return &P1{}
}

// PrintP1 p1实现i1接口
func (p *P1) PrintP1() {
	fmt.Println("Hello, This is P1")
}

func (p *P1) DoSomething() {
	// 直接调用p2的方法
	p22 := p2.NewP2(p)
	p22.PrintP2()
}
