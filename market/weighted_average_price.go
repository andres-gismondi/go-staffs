package market

type WeightedAveragePrice struct {
	value    float64
	quantity int
}

func NewWAP() *WeightedAveragePrice {
	return &WeightedAveragePrice{
		value:    0,
		quantity: 0,
	}
}

func (wap *WeightedAveragePrice) addPurchase(value float64, quantity int) {
	total := (wap.value * float64(wap.quantity)) + (value * float64(quantity))
	wap.value = total / float64(wap.quantity+quantity)
	wap.quantity += quantity
}

func (wap *WeightedAveragePrice) updateQuantity(quantity int) {
	wap.quantity -= quantity
}
