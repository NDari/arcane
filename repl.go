package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	ns := TopLevel()
	rd := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("incantation ::= ")
		input, _ := rd.ReadString('\n')
		if strings.TrimSpace(input) == ":q" {
			fmt.Println("Goodbye!")
			break
		}
		s := Str{input}
		anys, err := Read(s)
		if err != nil {
			fmt.Println(fmt.Sprintf("reader failed:\n%v", err))
			continue
		}
		for _, a := range anys {
			v, err := Eval(ns, a)
			if err != nil {
				fmt.Println(fmt.Sprintf("eval failed:\n%v", err))
				continue
			}
			fmt.Println(*v)
		}
	}
}
