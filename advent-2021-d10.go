package main

import "fmt"

func main() {

	opened := []string{}
	errors := map[string]int{}

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
