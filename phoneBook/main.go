package main

import (
	"fmt"
	"os"
	"path"
)

type Entry struct {
	Name    string
	Surname string
	Tel     string
}

var data []Entry

func search(key string) *Entry {
	for _, v := range data {
		if v.Surname == key {
			return &v
		}
	}
	return nil
}

func list() {
	for _, v := range data {
		fmt.Println(v)
	}
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		exe := path.Base(arguments[0])
		fmt.Printf("Usage: %s search|list <arguments>\n", exe)
		return
	}

	data = append(data, Entry{"Almaz", "Yessenyazov", "87761688760"})
	data = append(data, Entry{"John", "Doe", "8776112312312"})
	data = append(data, Entry{"Jane", "Jackson", "1324332343421"})

	switch arguments[1] {
	case "search":
		if len(arguments) != 3 {
			exe := path.Base(arguments[0])
			fmt.Printf("Usage: %s search <Surname>\n", exe)
			return
		}
		result := search(arguments[2])
		if result == nil {
			fmt.Println("Entry not found:", arguments[2])
			return
		}
		fmt.Println(*result)
	case "list":
		list()
	default:
		fmt.Println("Not a valid option")
	}
}
