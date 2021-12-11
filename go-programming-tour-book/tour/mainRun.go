package main

import (
	"log"
	"tour/cmd"
)

func main() {
	//flagExp1()
	//CLIExp()

	// =======part2
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}