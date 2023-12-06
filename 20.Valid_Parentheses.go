package leetcode

import "errors"

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	st := newStack(len(s) / 2)
	for i := range s {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			if err := st.put(s[i]); err != nil {
				return false
			}
			continue
		}

		if (s[i] - st.pop()) > 2 {
			return false
		}
	}
	return st.ptr == 0
}

type stack struct {
	buf []byte
	ptr int
}

func newStack(len int) stack {
	return stack{
		buf: make([]byte, len),
	}
}

func (s *stack) put(in byte) error {
	if len(s.buf)-1 < s.ptr {
		return errors.New("stack over flow")
	}
	s.buf[s.ptr] = in
	s.ptr++
	return nil
}

func (s *stack) pop() byte {
	if s.ptr == 0 {
		return 0
	}
	s.ptr--
	return s.buf[s.ptr]
}
