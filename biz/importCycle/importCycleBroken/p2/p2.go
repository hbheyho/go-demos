package p2

import (
	"fmt"
)

// 定义一个p1实现的接口i1
type i1 interface {
	PrintP1()
}

// P2 p2持有i1接口
type P2 struct {
	I1 i1
}

func NewP2(I1 i1) *P2 {
	return &P2{
		I1: I1,
	}
}

func (p *P2) PrintP2() {
	fmt.Println("Hello, This is P2")
}

func (p *P2) DoSomething() {
	// 通过接口调用p1的方法, 而不是直接引用
	p.I1.PrintP1()
}
