package indicator

import (
	"testing"
)

func TestAtr(t *testing.T) {
	trs := []*TR{
		{TR: 1},
		{TR: 3},
		{TR: 11},
		{TR: 13},
		{TR: -6},
		{TR: 18.05},
	}

	want := []*ATR{
		{ATR: 0},
		{ATR: 0},
		{ATR: 5},
		{ATR: 9},
		{ATR: 6},
		{ATR: 8.35},
	}

	got := Atr(trs, 3)

	for index, atr := range got {
		if atr.ATR != want[index].ATR {
			t.Errorf("ATR index:%d: want:%v got:%v", index, want[index].ATR, atr.ATR)
		}
	}
}
