package indicator

type Peroid string

const (
	P1M Peroid = "1m"
	P1D Peroid = "1d"
)

func (p Peroid) Seconds() int64 {
	switch p {
	case P1M:
		return 60
	case P1D:
		return 60 * 60 * 24
	default:
		return 0
	}
}
