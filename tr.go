package indicator

import "math"

type TR struct {
	TR float64
}

func Tr(quotes []*Quote) []*TR {
	trs := make([]*TR, 0, len(quotes))
	var tr float64
	for index, quote := range quotes {
		tr = quote.High - quote.Low
		if index > 0 {
			tr = math.Max(math.Max(tr, math.Abs(quotes[index-1].Close-quote.High)), math.Abs(quotes[index-1].Close-quote.Low))
		}

		trs = append(trs, &TR{TR: tr})
	}

	return trs
}
