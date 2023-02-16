package main

import (
	"bufio"
	f "fmt"
	"log"
	"os"
)

// This is a shorthand for fmt.PrintLn
var pl = f.Println

/*
The following one is a go function
this is a multiline comment in go
*/

func main() {
	pl("Hello Go!")
	pl("What is your name?")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err == nil {
		pl("Hello", name)
	} else {
		log.Fatal(err)
	}

}
