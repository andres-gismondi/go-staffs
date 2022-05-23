package value_object

import "fmt"

// Value Objects are immutable. There is no single cause, reason, or other argument
// to change the state of Value Object during its lifetime. Whenever we want to change
// an inner state of VO or combine multiple of them, we always need to return a new instance

// wrong way to change the state
func (m *Money) AddMoney(amount int) {
	m.Value += amount
}

//right way
func (m Money) WithAmount(amount int) Money {
	return Money{
		Value:    m.Value + amount,
		Currency: m.Currency,
	}
}

// Good practice in Golang is to bind functions to values instead of references of VO
func (m Money) Deduct(other Money) (Money, error) {
	if !m.Currency.EqualTo(other.Currency) {
		return Money{}, fmt.Errorf("currencies must be identical")
	}

	if m.Value < other.Value {
		return Money{}, fmt.Errorf("there is not enought value to deduct")
	}

	return Money{
		Value:    m.Value - other.Value,
		Currency: m.Currency,
	}, nil
}

// We should not validate VO during its whole lifetime, but only in creation
// when we want to create a new VO, we must always perform validation and return errors.
