package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	files := make(map[string]string)
	
	files["hello.txt"] = "Hello, modoki-OS world!"

	fmt.Println("started modoki-OS v2.0")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("$: ")

		if !scanner.Scan() {
			break
		}

		line := scanner.Text()
		words := strings.Fields(line)

		if len(words) == 0 {
			continue
		}

		cmd := words[0]

		switch cmd {
		case "ls":
			if len(files) == 0 {
				fmt.Println("(empty)")
			}
			for fileName := range files {
				fmt.Println(fileName)
			}

		case "touch":
			if len(words) < 2 {
				fmt.Println("usage: touch <filename> [content]")
				continue
			}
			name := words[1]
			content := ""
			if len(words) > 2 {
				content = strings.Join(words[2:], " ")
			}
			files[name] = content
			fmt.Printf("File '%s' created.\n", name)

		case "cat":
			if len(words) < 2 {
				fmt.Println("usage: cat <filename>")
				continue
			}
			name := words[1]
			if content, ok := files[name]; ok {
				fmt.Println(content)
			} else {
				fmt.Printf("cat: %s: No such file\n", name)
			}

		case "rm":
			if len(words) < 2 {
				fmt.Println("usage: rm <filename>")
				continue
			}
			delete(files, words[1])

		case "exit":
			fmt.Println("bye!")
			return

		default:
			fmt.Println("command not found:", cmd)
		}
	}
}
