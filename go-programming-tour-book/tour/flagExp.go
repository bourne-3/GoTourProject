package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
)

type Name string

func (n *Name) String() string {
	return fmt.Sprint(*n)
}

func (i *Name) Set(value string) error {
	if len(*i) > 0 {
		return errors.New("name flag already set")
	}

	*i = Name("There is a custom Name:" + value)
	return nil
}

func RunCustom() {
	var name Name
	flag.Var(&name, "name", "帮助信息")
	flag.Parse()

	log.Printf("name:%s", name)
}





func flagExp1() {
	var name string
	// 有两个值的话，会使用后面那个
	// 会直接把值就赋给了name了
	flag.StringVar(&name, "name", "DefalutVal", "Prompt Info name")
	flag.StringVar(&name, "n", "DefalutVal", "Prompt Info n")


	var gender string
	flag.StringVar(&gender, "gender", "DefaultVal", "gender for this person")

	flag.Parse()
	log.Println(name, gender)
}


func CLIExp() {
	flag.Parse()
	args := flag.Args()

	if (len(args) <= 0) {  // Args最少是几个？  最少是0，args[0]不是文件名，是第一个参数，区别与os.Args
		return
	}
	var name string

	switch args[0] {
	case "go":
		goCmd := flag.NewFlagSet("go", flag.ExitOnError)
		goCmd.StringVar(&name, "name", "goDefaultVal", "prompt")
		_ = goCmd.Parse(args[1:])
	case "php":
		goCmd := flag.NewFlagSet("php", flag.ExitOnError)
		goCmd.StringVar(&name, "name", "phpDefaultVal", "prompt")
		_ = goCmd.Parse(args[1:])
	case "java":
		// 子命令
		javaCmd := flag.NewFlagSet("java", flag.ExitOnError)
		javaCmd.StringVar(&name, "name", "javaDefaultVal", "tishi")
		_ = javaCmd.Parse(args[1:])
	}
	log.Printf("name : %s", name)
}

func CustomFlag() {
	// flag包中的Value

}
