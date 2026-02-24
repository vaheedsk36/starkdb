package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const logFileName = "data.log"

func main() {
	store := make(map[string]string)

	// Open log file
	file, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Recover state from log
	fmt.Println("Recovering from log...")
	recoverFromLog(file, store)

	fmt.Println("MiniDB started. Type EXIT to quit.")
	startREPL(store, file)
}

func recoverFromLog(file *os.File, store map[string]string) {
	// Move cursor to beginning
	file.Seek(0, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		applyCommand(line, store, false)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error during recovery:", err)
	}

	// Move cursor back to end for appending
	file.Seek(0, os.SEEK_END)
}

func startREPL(store map[string]string, file *os.File) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()

		if strings.ToUpper(input) == "EXIT" {
			fmt.Println("Bye 👋")
			return
		}

		applyCommand(input, store, true, file)
	}
}

func applyCommand(input string, store map[string]string, logWrite bool, file ...*os.File) {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return
	}

	command := strings.ToUpper(parts[0])

	switch command {

	case "SET":
		if len(parts) != 3 {
			fmt.Println("Usage: SET key value")
			return
		}
		store[parts[1]] = parts[2]
		fmt.Println("OK")

		if logWrite {
			appendToLog(file[0], input)
		}

	case "GET":
		if len(parts) != 2 {
			fmt.Println("Usage: GET key")
			return
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
			return
		}
		delete(store, parts[1])
		fmt.Println("Deleted")

		if logWrite {
			appendToLog(file[0], input)
		}

	case "KEYS":
		if len(store) == 0 {
			fmt.Println("(empty)")
			return
		}
		for k := range store {
			fmt.Println(k)
		}

	default:
		fmt.Println("Unknown command")
	}
}

func appendToLog(file *os.File, input string) {
	_, err := fmt.Fprintln(file, input)
	if err != nil {
		fmt.Println("Log write error:", err)
		return
	}

	// Force disk flush (Durability)
	file.Sync()
}