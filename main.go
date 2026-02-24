package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	store := make(map[string]string)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("MiniDB started. Type EXIT to quit.")

	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()

		if strings.ToUpper(input) == "EXIT" {
			break
		}

		parts := strings.Fields(input)

		if len(parts) == 0 {
			continue
		}

		command := strings.ToUpper(parts[0])

		switch command {
		case "SET":
			if len(parts) != 3 {
				fmt.Println("Usage: SET key value")
				continue
			}
			store[parts[1]] = parts[2]
			fmt.Println("OK")

		case "GET":
			if len(parts) != 2 {
				fmt.Println("Usage: GET key")
				continue
			}
			value, exists := store[parts[1]]
			if !exists {
				fmt.Println("(nil)")
			} else {
				fmt.Println(value)
			}

		case "DELETE":
			if len(parts) != 2 {
				fmt.Println("Usage: DELETE key")
				continue
			}
			delete(store, parts[1])
			fmt.Println("Deleted")

		default:
			fmt.Println("Unknown command")
		}
	}
}