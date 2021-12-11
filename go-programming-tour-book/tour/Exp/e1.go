package main

import (
	"flag"
	"fmt"
	"strings"
	"tour/internal/word"
)

func t1() {
	flag.Parse()
	args := flag.Args()
	fmt.Println(args)
}
func t2()  {
	// 测试iota
	const (
		mon int = iota
		tue
		wed
	)
	fmt.Println(wed)  // 2
}

func t3() {
	// 测试大驼峰
	s := "number_class_salary"
	s = word.UnderscoreToUpperCamelCase(s)
	fmt.Println(s)

	s2 := "RunCom"
	s2 = word.CamelCaseToUnderscore(s2)
	fmt.Println(s2)
}

func t4() {
	// 测试join
	s := strings.Join([]string{"广东", "外语", "外贸"}, "\n")
	fmt.Println(s)

}