package props

import "fmt"

type (
	Measurement string
	Unit        struct {
		Size        float64
		Measurement Measurement
	}
)

const (
	Px      Measurement = "px"
	Percent Measurement = "%"
	Rem     Measurement = "rem"
)

func (u Unit) String() string {
	switch u.Measurement {
	case Px:
		return fmt.Sprintf("%.0f%s", u.Size, u.Measurement)
	case Percent:
		return fmt.Sprintf("%.2f%s", u.Size, u.Measurement)
	case Rem:
		return fmt.Sprintf("%.3f%s", u.Size, u.Measurement)
	}
	return ""
}
