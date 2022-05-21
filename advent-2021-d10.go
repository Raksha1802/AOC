package main

import "fmt"

func main() {
	opened := []string{}
	errors := []string{}
	//input := inputs[0]

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
				errors = append(errors, currentChar)
				break
			}
		}
	}

	fmt.Println()
	fmt.Println("length", errors)

	var score int
	for i := range errors {
		fmt.Println()
		fmt.Println(errors[i])
		score += scores[errors[i]]
		fmt.Printf("\nScore %d", score)
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
