package stack

import (
	"bytes"
	"fmt"
)

type element [4]int

func (e *element) Pop() int {
	if len(e) > 0 {
		return (*e)[0]
	}

	return int(nil)
}

func (e *element) Push(i int) bool {
	if len(e) >= 4 {
		return false
	}

	e[len(e)-1] = i
	return true
}

func (e *element) String() string {
	bt := bytes.NewBufferString("")
	for i, index := range e {
		bt.WriteString("[")
		bt.WriteString(fmt.Sprint(index))
		bt.WriteString(":")
		bt.WriteString(fmt.Sprint(i))
		bt.WriteString("]")
	}
	return bt.String()
}

func (e *element) newStackArr() *element {
	return &element{}
}
