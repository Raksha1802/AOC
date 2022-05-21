# -*- coding: utf-8 -*-
"""advent-2021-d10.ipynb
"""

f = open("input.txt", "r")
inputs = f.read().splitlines()
# print(inputs)

scores = {")": 3, "]": 57, "}": 1197, ">": 25137}
openers = {"(": ")", "[": "]", "{": "}", "<": ">"}
closers = { ")": "(", "]": "[", "}": "{", ">": "<"}

opened = list()
errors = list()
for input in inputs:
  for i in range(len(input)):
    currentChar = input[i]
    if currentChar in openers:
      opened.append(currentChar)
      continue
    elif currentChar in closers:
      if opened[len(opened)-1] == closers[currentChar]:
        opened = opened[:len(opened)-1]
        continue
      else:
        print(currentChar)
        # errors.append(currentChar)
        break

score = 0
for i in range(len(errors)):
  score += scores[errors[i]]
print("Total Syntax Scoring:", score)