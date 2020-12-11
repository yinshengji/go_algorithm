package stack

import "container/list"

type Stack struct {
	data *list.List
}

//NewStack 初始化栈
func NewStack() *Stack {
	return &Stack{list.New()}
}

//Push 入栈
func (stack *Stack) Push(value interface{}) {
	stack.data.PushBack(value)
}

//Pop 出栈
func (stack *Stack) Pop() interface{} {
	e := stack.data.Back()
	if e != nil {
		stack.data.Remove(e)
		return e.Value
	}
	return nil
}

//Peak 获取栈顶
func (stack *Stack) Peak() interface{} {
	e := stack.data.Back()
	if e != nil {
		return e.Value
	}

	return nil
}

//Len 获取栈长度
func (stack *Stack) Len() int {
	return stack.data.Len()
}

//Empty 判断栈是否为空
func (stack *Stack) Empty() bool {
	return stack.data.Len() == 0
}
