package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const logFileName = "data.log"
const snapshotFileName = "snapshot.db"

func main() {
	store := make(map[string]string)

	logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	fmt.Println("Loading snapshot...")
	loadSnapshot(store)

	fmt.Println("Recovering from log...")
	recoverFromLog(logFile, store)

	fmt.Println("MiniDB started. Type EXIT to quit.")
	startREPL(store, logFile)
}

func loadSnapshot(store map[string]string) {
	file, err := os.Open(snapshotFileName)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			store[parts[0]] = parts[1]
		}
	}
}

func recoverFromLog(file *os.File, store map[string]string) {
	file.Seek(0, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		applyCommand(scanner.Text(), store, false, file)
	}

	file.Seek(0, os.SEEK_END)
}

func startREPL(store map[string]string, logFile *os.File) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()

		if strings.ToUpper(input) == "EXIT" {
			fmt.Println("Bye 👋")
			return
		}

		applyCommand(input, store, true, logFile)
	}
}

func applyCommand(input string, store map[string]string, logWrite bool, logFile *os.File) {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return
	}

	cmd := strings.ToUpper(parts[0])

	switch cmd {

	case "SET":
		if len(parts) != 3 {
			fmt.Println("Usage: SET key value")
			return
		}
		store[parts[1]] = parts[2]
		fmt.Println("OK")
		if logWrite {
			appendToLog(logFile, input)
		}

	case "GET":
		if len(parts) != 2 {
			fmt.Println("Usage: GET key")
			return
		}
		val, ok := store[parts[1]]
		if !ok {
			fmt.Println("(nil)")
		} else {
			fmt.Println(val)
		}

	case "DELETE":
		if len(parts) != 2 {
			fmt.Println("Usage: DELETE key")
			return
		}
		delete(store, parts[1])
		fmt.Println("Deleted")
		if logWrite {
			appendToLog(logFile, input)
		}

	case "KEYS":
		if len(store) == 0 {
			fmt.Println("(empty)")
			return
		}
		for k := range store {
			fmt.Println(k)
		}

	case "SNAPSHOT":
		createSnapshot(store)
		clearLog(logFile)
		fmt.Println("Snapshot created and log cleared.")

	default:
		fmt.Println("Unknown command")
	}
}

func appendToLog(file *os.File, input string) {
	fmt.Fprintln(file, input)
	file.Sync()
}

func createSnapshot(store map[string]string) {
	file, err := os.Create(snapshotFileName)
	if err != nil {
		fmt.Println("Snapshot error:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for k, v := range store {
		fmt.Fprintf(writer, "%s=%s\n", k, v)
	}
	writer.Flush()
	file.Sync()
}

func clearLog(file *os.File) {
	file.Truncate(0)
	file.Seek(0, 0)
}