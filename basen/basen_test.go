package basen

import (
	"testing"
)

func TestNewBaseN(t *testing.T) {
	dn := NewBaseN(Base62Table)
	if dn == nil {
		t.Error("Unexpected behaviour")
	}
}

func TestEncode(t *testing.T) {
	dn := NewBaseN(Base62Table)
	if dn.Encode(0) != "0" {
		t.Error("Unexpected behaviour")
	}
	if dn.Encode(61) != "Z" {
		t.Error("Unexpected behaviour")
	}
	if dn.Encode(62) != "10" {
		t.Error("Unexpected behaviour")
	}
}

func TestDecode(t *testing.T) {
	dn := NewBaseN(Base62Table)
	if dn.Decode("0") != 0 {
		t.Error("Unexpected behaviour")
	}
	if dn.Decode("Z") != 61 {
		t.Errorf("Unexpected behaviour %d", dn.Decode("Z"))
	}
	if dn.Decode("10") != 62 {
		t.Errorf("Unexpected behaviour %d", dn.Decode("10"))
	}
}

func TestEncodeDecode(t *testing.T) {
	dn := NewBaseN(Base62Table)

	for i := 0; i < 10; i++ {
		if dn.Decode(dn.Encode(uint(i))) != i {
			t.Error("Unexpected behaviour")
		}
	}
}
