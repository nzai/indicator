package indicator

import (
	"testing"
)

func TestTr(t *testing.T) {
	quotes := []*Quote{
		{High: 11, Low: 10, Close: 10},
		{High: 12, Low: 9, Close: 9.5},
		{High: 15, Low: 10, Close: 15},
		{High: 10, Low: 5, Close: 8},
		{High: 7, Low: 4.3, Close: 5.1},
	}

	want := []*TR{
		{TR: 1},
		{TR: 3},
		{TR: 5.5},
		{TR: 10},
		{TR: 3.7},
	}

	got := Tr(quotes)

	for index, tr := range got {
		if tr.TR != want[index].TR {
			t.Errorf("TR index:%d: want:%v got:%v", index, want[index].TR, tr.TR)
		}
	}
}
