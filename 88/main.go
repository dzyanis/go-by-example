package main

import (
	"fmt"
        "flag"
)

type T struct {
	nums []int
	len  int
}

func NewT(len int) *T {
	sl := make([]int, len)
	for i := 0; i < len; i++ {
		sl[i] = 1
	}
	return &T{
		nums: sl,
		len:  len,
	}
}

func (t *T) Incr() bool {
	for i := t.len - 1; i >= 0; i-- {
		t.nums[i]++
		if t.nums[i] > 9 && i == 0 {
			return false
		} else if t.nums[i] > 9 {
			t.nums[i] = 1
		} else {
			return true
		}

	}

	return false
}

func (t *T) M() int {
	m := 1
	for _, n := range t.nums {
		m *= n
	}
	return m
}

func (t *T) S() int {
	s := 0
	for _, n := range t.nums {
		s += n
	}
	return s
}

func (t *T) String() string {
	return fmt.Sprintf("%#v, s:%d, m:%d", t.nums, t.M(), t.S())
}

func main() {
        var l = flag.Int("len", 2, "length")
        flag.Parse()
	b := NewT(*l)
	for b.Incr() {
		if b.M() == b.S() {
			fmt.Println(b)
			break
		}
	}
	fmt.Println("finish")
}

