package indicator

type ATR struct {
	ATR float64
}

func Atr(trs []*TR, peroid int) []*ATR {
	if peroid <= 0 || len(trs) < peroid {
		return []*ATR{}
	}

	_trs := make([]float64, len(trs))
	for index, tr := range trs {
		_trs[index] = tr.TR
	}

	smas := Sma(_trs, peroid)

	atrs := make([]*ATR, len(trs))
	for index, sma := range smas {
		atrs[index] = &ATR{ATR: sma}
	}

	return atrs
}
