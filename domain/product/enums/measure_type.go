package enums

type MeasureType int

const (
	Thing MeasureType = iota
	Liter
	Meter
)

func IsValid(measure MeasureType) bool {
	return measure >= Thing && measure <= Meter
}
