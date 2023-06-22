package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func init() {
	fmt.Println("init çalıştı")
}

var variable string = get()

func get() string {
	fmt.Println("get çalıştı")
	return "Emre"
}

func main() {
	var text string = "Hello " + os.Args[1]
	fmt.Println(text)

	fp, _ := os.Create("emre.txt")
	// fp.WriteString(text)
	io.Copy(fp, strings.NewReader(text))
	defer fp.Close()
}
