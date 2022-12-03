package main

import "fmt"

// Type Assertion 作用：
// 	1. 检查 i 是否为 nil
//	2. 检查 i 存储的值是否为某个类型
func main() {

	// use 1: t := i.(T) => 接口对象 i 存储的值的类型是 T, 如果断言成功，就会返回值给 t，如果断言失败，就会触发 panic
	var i1 interface{} = 10
	v1 := i1.(int)
	fmt.Println(v1)

	// panic: interface {} is int, not string
	s1 := i1.(string)
	fmt.Println(s1)

	// value is nil
	// panic:  interface is nil, not interface {}
	var i2 interface{} // nil
	_ = i2.(interface{})

	// use 2: t, ok:= i.(T) => 不直接返回panic, 而是通过返回的bool来判断是否断言成功
	var i3 interface{} = 10
	v3, ok := i3.(int)
	if ok {
		fmt.Printf("Assertion success: %d \n", v3)
	}

	s3, ok := i3.(string)
	if !ok {
		fmt.Printf("Assertion fail: %s", s3)
	}

	// use 3: type switch => 需要区分多种类型，可以使用 type switch 断言，这个将会比一个一个进行类型断言更简单、直接、高效
	findType(10)
	findType("10")
	var i4 interface{}
	findType(i4)

}

func findType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println(i, "is int")
	case string:
		fmt.Println(i, "is string")
	case nil:
		fmt.Println(i, "is nil")
	default:
		fmt.Println(i, "not type matched")
	}
}
