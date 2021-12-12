package main

import "fmt"

// 测试iota
const (
	No1 = 1 + iota
	No2
	No3
)

type Bourne struct {
	name string
	age  int
}

func main() {
	t1()
}

func t1() {
	fmt.Println(No1, No3)
	b := Bourne{name: "bourne", age: 18}
	fmt.Println(b)
}

func (b Bourne) String() string {
	return fmt.Sprintf("使用String()进行重写，我的名字叫做%s，我的年龄为:%d", b.name, b.age)
}
