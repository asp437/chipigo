package main

import "testing"

func TestStack(t *testing.T) {
	stack := new(Stack)
	stack.init(16)
	a, b, c, d, e, f := uint32(15), uint32(42), uint32(32), uint32(52), uint32(435), uint32(224)

	stack.push(a)
	stack.push(b)
	stack.push(c)
	stack.push(d)
	stack.push(e)
	stack.push(f)
	ft := stack.pop()
	et := stack.pop()
	dt := stack.pop()
	ct := stack.pop()
	bt := stack.pop()
	at := stack.pop()
	if ft != f {
		t.Errorf("Stack.pop return <%d>, <%d> expected", ft, f)
	}
	if et != e {
		t.Errorf("Stack.pop return <%d>, <%d> expected", et, e)
	}
	if dt != d {
		t.Errorf("Stack.pop return <%d>, <%d> expected", dt, d)
	}
	if ct != c {
		t.Errorf("Stack.pop return <%d>, <%d> expected", ct, c)
	}
	if bt != b {
		t.Errorf("Stack.pop return <%d>, <%d> expected", bt, b)
	}
	if at != a {
		t.Errorf("Stack.pop return <%d>, <%d> expected", at, a)
	}
}
