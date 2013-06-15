package main

import "testing"

func TestStack(t *testing.T) {
	stack := new(Stack)
	stack.init(16)
	a, b := uint32(15), uint32(42)

	stack.push(a)
	stack.push(b)
	bt := stack.pop()
	at := stack.pop()
	if bt != b {
		t.Errorf("Stack.pop return <%d>, <%d> expected", bt, b)
	}
	if at != a {
		t.Errorf("Stack.pop return <%d>, <%d> expected", at, a)
	}
}
