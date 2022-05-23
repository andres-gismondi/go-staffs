package value_object

type Color struct {
	Red   float64
	Green float64
	Blue  float64
}

// we need to check all fields equality
func (c Color) EqualTo(other Color) bool {
	return c.Blue == other.Blue && c.Red == other.Red && c.Green == other.Green
}

type Currency struct {
	ID string
}
type Money struct {
	Value int
	Currency
}

func (c Currency) EqualTo(other Currency) bool {
	return c.ID == other.ID
}
func (m Money) EqualTo(other Money) bool {
	return m.Value == other.Value && m.Currency.EqualTo(other.Currency)
}
