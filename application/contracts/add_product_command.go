package contracts

type AddProductCommand struct {
	Name        string
	Measure     int32
	Code        string
	Description string
}
