package main

import (
	"fmt"
	"regexp"
)

func main() {
	cmd := "echo \"Hello world\"!\x00"
	re := regexp.MustCompile(`("[^"]+?"|\S+)`)
	args := re.FindAllString(cmd, -1)
	fmt.Println("%v", args)
	fmt.Println("%v", len(args))
}
