
// the idea is store min value each round and encoded the actual value 
// corresponding to the min
type MinStack struct {
	min int
	stack []int
}

func Constructor() MinStack {
	return MinStack{
		min: math.MaxInt64,
		stack: []int{},
	}
}

func (m *MinStack) Push(val int) {
	if len(m.stack) == 0 {
		m.stack = append(m.stack, 0)
		m.min = val
	} else {
		m.stack = append(m.stack, val - m.min)
		if val < m.min{
			m.min = val
		}
	}
}

func (m *MinStack) Pop() {
	if len(m.stack) == 0{
		return
	}
	pop := m.stack[len(m.stack)-1]
	m.stack = m.stack[:len(m.stack)-1]
	if pop <0{
		m.min -= pop
	}
}

func (m *MinStack) Top() int {
	top := m.stack[len(m.stack)-1]
	if top > 0{
		return top + m.min
	} else {
		return m.min
	}
}

func (m *MinStack) GetMin() int {
	return m.min
}
