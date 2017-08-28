package basen

import (
	"fmt"
	"math"
)

const (
	Base62Table = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	BaseBinaryTable = "01"
	BaseHexTable = "0123456789ABCDEF"
)

type BaseN struct {
	table string
	hash map[rune]int
	n uint
}

func NewBaseN(table string) *BaseN {
	bn := &BaseN{
		table: table,
		hash: make(map[rune]int, len(table)),
		n: uint(len(table)),
	}

	for i, r := range table {
		bn.hash[r] = i
	}

	return bn
}

func (bn *BaseN) Encode(n uint) string {
	b := ""
	d := n
	m := uint(0)
	for d >= bn.n {
		m = d % bn.n
		b = fmt.Sprintf("%s%s", string(bn.table[m]), b)
		d = d / bn.n
	}
	if d > 0 || len(b) == 0 {
		b = fmt.Sprintf("%s%s", string(bn.table[d]), b)
	}
	return b
}

func (bn *BaseN) Decode(b string) int {
	n := 0
	len := len(b) - 1
	for i, r := range b {
		v := bn.hash[r]
		n += v * int(math.Pow(float64(bn.n), float64(len - i)))
	}
	return n
}
