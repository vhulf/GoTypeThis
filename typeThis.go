package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"unicode"

	"github.com/go-vgo/robotgo"
)

const (
	delayBeforeTyping = 10 * time.Second // Adjust the delay as needed
	delayBetweenRunes = 20               // adjust typing speed as needed (ms)
)

func returnRelatedKeyForSpecialUpper(char string) string {
	switch char {
	case string("'"):
		return string("'")
	case "~":
		return "`"
	case "!":
		return "1"
	case "@":
		return "2"
	case "#":
		return "3"
	case "$":
		return "4"
	case "%":
		return "5"
	case "^":
		return "6"
	case "&":
		return "7"
	case "*":
		return "8"
	case "(":
		return "9"
	case ")":
		return "0"
	case "_":
		return "-"
	case "+":
		return "="
	case "{":
		return "["
	case "}":
		return "]"
	case "|":
		return "\\"
	case ":":
		return ";"
	case "<":
		return ","
	case ">":
		return "."
	case "?":
		return "/"
	}

	return "FAILED"
}

func typeFileContents(filePath string) error {
	specialUppers := "~!@#$%^&*()_+{}:<>?|"
	specialUppers += string('"')

	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	time.Sleep(delayBeforeTyping)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() { // Type the line followed by a newline
		line := scanner.Text() + "\n"
		for i := 0; i < len(line); i++ {
			tmpchar := string(line[i])

			if unicode.IsUpper([]rune(tmpchar)[0]) {
				robotgo.KeyDown("rshift", "down")
				robotgo.KeyTap(string(tmpchar))
				robotgo.KeyUp("rshift", "up")
			} else if strings.ContainsAny(tmpchar, specialUppers) { //check if special upper!
				tmpchar = returnRelatedKeyForSpecialUpper(string(tmpchar))
				robotgo.KeyDown("rshift", "down")
				robotgo.KeyTap(string(tmpchar))
				robotgo.KeyUp("rshift", "up")
			} else {
				robotgo.KeyTap(string(tmpchar))
			}

		}

		robotgo.KeyTap(string(robotgo.Enter))
		time.Sleep(delayBetweenRunes)

	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <file_path>")
		return
	}

	filePath := os.Args[1]
	err := typeFileContents(filePath)
	fmt.Println("Typed!")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
