package numtext

import "fmt"

type Magnitude struct {
	Magnitude int64
	Symbol    string
}

func (m *Magnitude) ToText(number int64) string {
	return fmt.Sprintf("%.1f%s", float64(number)/float64(m.Magnitude), m.Symbol)
}

func ToText(number int64) string {
	list := []Magnitude{
		Magnitude{1000000000000, "T"},
		Magnitude{1000000000, "B"},
		Magnitude{1000000, "M"},
	}
	for _, m := range list {
		if m.Magnitude < int64(number) {
			return m.ToText(int64(number))
		}
	}
	return fmt.Sprintf("%d", number)

}
