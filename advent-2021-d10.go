package main

import "fmt"

func main() {
	opened := []string{}
	errors := []string{}
	input := inputs[2]

	for i := 0; i < len(input); i++ {

		currentChar := string(input[i])

		if _, isOpener := openers[currentChar]; isOpener {
			opened = append(opened, currentChar)
			continue
		} else if _, isCloser := closers[currentChar]; !isCloser {
			continue
		}

		fmt.Printf("\ncurrentChar: %s", currentChar)
		fmt.Printf("\nopened[len(opened)-1]: %s", opened[len(opened)-1])

		if opened[len(opened)-1] == closers[currentChar] {
			opened = opened[:len(opened)-2]
			continue
		} else {
			errors = append(errors, currentChar)
		}
	}
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

var inputs = []string{
	"[({(<(())[]>[[{[]{<()<>>",
	"[(()[<>])]({[<{<<[]>>(",
	"{([(<{}[<>[]}>{[]{[(<()>",
	"(((({<>}<{<{<>}{[]{[]{}",
	"[[<[([]))<([[{}[[()]]]",
	"[{[{({}]{}}([{[{{{}}([]",
	"{<[[]]>}<{[{[{[]{()[[[]",
	"[<(<(<(<{}))><([]([]()",
	"<{([([[(<>()){}]>(<<{{",
	"<{([{{}}[<[[[<>{}]]]>[]]",
}
