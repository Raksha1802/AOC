package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	opened := []string{}
	errors := map[string]int{}
	inputs := ReadFile() // remove for hardcoded input

	for _, input := range inputs {
		for i := 0; i < len(input); i++ {
			currentChar := string(input[i])
			if _, isOpener := openers[currentChar]; isOpener {
				opened = append(opened, currentChar)
				continue
			} else if _, isCloser := closers[currentChar]; !isCloser {
				continue
			}
			if opened[len(opened)-1] == closers[currentChar] {
				opened = opened[:len(opened)-1]
				continue
			} else {
				errors[currentChar]++
				break
			}
		}
	}
	var total int = 0
	for err, times := range errors {
		fmt.Printf("\n %s: %d", err, times)
		total += scores[err] * times
	}
	fmt.Printf("\ntotal: %d", total)
}

var closers = map[string]string{
	")": "(",
	"]": "[",
	"}": "{",
	">": "<",
}

var openers = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
	"<": ">",
}

var scores = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}

// Input hardcoded

// var inputs = []string{
// 	"[({(<(())[]>[[{[]{<()<>>",
// 	"[(()[<>])]({[<{<<[]>>(",
// 	"{([(<{}[<>[]}>{[]{[(<()>",
// 	"(((({<>}<{<{<>}{[]{[]{}",
// 	"[[<[([]))<([[{}[[()]]]",
// 	"[{[{({}]{}}([{[{{{}}([]",
// 	"{<[[]]>}<{[{[{[]{()[[[]",
// 	"[<(<(<(<{}))><([]([]()",
// 	"<{([([[(<>()){}]>(<<{{",
// 	"<{([{{}}[<[[[<>{}]]]>[]]",
// }

// Read input from text file

func ReadFile() []string {

	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	lines := []string{}
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	defer readFile.Close()
	return lines
}
