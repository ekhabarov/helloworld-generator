package main

import "fmt"

var tpl = `package main

import (
	"flag"
	"fmt"
)

var name = flag.String("name", "World", "Just a name")

func main() {
	flag.Parse()

	fmt.Printf("Hello, %s!\n", *name)
}`

func main() {
	fmt.Println(tpl)
}
