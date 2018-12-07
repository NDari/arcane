package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("incantation ::= ")
		str, _ := rd.ReadString('\n')
		if strings.TrimSpace(str) == ":q" {
			fmt.Println("Goodbye!")
			break
		}
		anys, err := Read(str)
		if err != nil {
			fmt.Println("reader failed:", err)
			continue
		}
		fmt.Println(anys.Repr())
	}
}
