package store

import (
	"testing"
)

func sliceEqual(x, y []interface{}) bool {
	if len(x) != len(y) {
		return false
	}
	for i, v := range x {
		if y[i] != v {
			return false
		}
	}
	return true
}

func TestAppend(t *testing.T) {
	cb := New(3)
	cb.Append(7)
	if cb.buf[0] != 7 {
		t.Error("did not append value")
	}
	if cb.tail != 1 {
		t.Error("did not increment tail")
	}
}

func TestWraparound(t *testing.T) {
	cb := New(3)
	cb.Append(1)
	cb.Append(2)
	cb.Append(3)
	if cb.tail != 0 {
		t.Error("did not wrap tail")
	}
	if !cb.wrapped {
		t.Error("did not update wrapped flag")
	}
	cb.Append(4)
	if cb.buf[0] != 4 {
		t.Error("did not append correctly")
	}
}

func TestSlices(t *testing.T) {
	cb := New(3)
	cb.Append(1)
	if !sliceEqual(cb.Slice(), []interface{}{1}) {
		t.Errorf("Bad slice; %v", cb.Slice())
	}
	cb.Append(2)
	if !sliceEqual(cb.Slice(), []interface{}{1, 2}) {
		t.Errorf("Bad slice; %v", cb.Slice())
	}
	cb.Append(3)
	if !sliceEqual(cb.Slice(), []interface{}{1, 2, 3}) {
		t.Errorf("Bad slice; %v", cb.Slice())
	}
	cb.Append(4)
	if !sliceEqual(cb.Slice(), []interface{}{2, 3, 4}) {
		t.Errorf("Bad slice; %v", cb.Slice())
	}
}
