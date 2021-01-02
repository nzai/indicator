package indicator

func Sma(values []float64, peroid int) []float64 {
	if peroid <= 0 || len(values) < peroid {
		return []float64{}
	}

	smas := make([]float64, len(values))
	var total float64
	for index, tr := range values {
		total += tr
		if index < peroid-1 {
			smas[index] = 0
			continue
		}

		if index >= peroid {
			total -= values[index-peroid]
		}

		smas[index] = total / float64(peroid)
	}

	return smas
}
