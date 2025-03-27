package main

// https://leetcode.com/problems/valid-parentheses/description/

type Stack struct {
	items []rune
}

func (s *Stack) Push(c rune) *Stack {
	s.items = append(s.items, c)
	return s
}

func (s *Stack) PushAll(c string) *Stack {
	for _, runeVal := range c {
		s.Push(runeVal)
	}
	return s
}

func (s *Stack) Pop() (rune, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	last := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return last, true
}

func (s Stack) IsEmpty() bool {
	return len(s.items) == 0
}

/*
func isValid(s string) bool {

	runeMap := map[rune]rune{'}': '{', ')': '(', ']': '['}

	// stack := new(Stack)
	stack := Stack{}

	for _, runeStr := range s {
		value, ok := runeMap[runeStr]

		if ok {
			stackedValue, _ := stack.Pop()

			if value != stackedValue {
				return false
			}

		} else {
			stack.Push(runeStr)
		}
	}

	return stack.IsEmpty()
}*/

func isValid(s string) bool {
	bracketMap := map[rune]rune{'}': '{', ')': '(', ']': '['}

	stack := Stack{}

	for _, bracketRune := range s {
		if openBracket, isClosing := bracketMap[bracketRune]; isClosing {
			if popped, IsEmpty := stack.Pop(); !IsEmpty || popped != openBracket {
				return false
			}
		} else {
			stack.Push(bracketRune)
		}
	}

	return stack.IsEmpty()

}

// func main() {

// 	s := "{()[{()}]}"

// 	fmt.Println(isValid(s))
// }
