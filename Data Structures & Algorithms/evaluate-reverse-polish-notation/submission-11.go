func evalRPN(tokens []string) int {
	s := Stack{}

	for _, token := range tokens{
		switch token{
			case "+":
				s.Push(s.Pop() + s.Pop())
			case "-":
				a := s.Pop()
				b:= s.Pop()
				s.Push(b-a)
			case "*":
				s.Push(s.Pop() * s.Pop())
			case "/":
			a := s.Pop()
				b:= s.Pop()
				s.Push(b / a)
			default:
				num, _ := strconv.Atoi(token)
				s.Push(num)
		}

	}
	return s.data[0]
}

type Stack struct{
	data []int
}

func Constructor() Stack{
	return Stack{[]int{}}
}

func (s *Stack) Push(val int){
	s.data= append(s.data, val)
}

func (s *Stack) Pop() int {
	val := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return val
}


