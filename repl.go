package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/NDari/arcane/reader"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	counter := 1
	for {
		fmt.Print(">> ")
		str, _ := rd.ReadString('\n')
		if strings.TrimSpace(str) == ":q" {
			fmt.Println("Goodbye!")
			break
		}
		r := reader.NewReader(str)
		anys, err := r.ReadAll()
		if err != nil {
			fmt.Println("read failed:\n", err)
			continue
		}
		fmt.Printf("$%d = ", counter)
		fmt.Println(anys)
		iter := anys.ToIterable()
		for {
			if iter.HasNext() {
				item := iter.Next()
				fmt.Print(item.Repr() + " ")
			} else {
				break
			}
		}
		fmt.Println()
		counter++
	}
}
