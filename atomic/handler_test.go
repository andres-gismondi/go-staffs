package atomic_test

import (
	"testing"
	"time"

	"go-staffs/atomic"
	"go-staffs/market"
)

func TestTimeSelect(t *testing.T) {
	timeSelect := atomic.NewTimeSelect()

	go func() {
		_ = timeSelect.Execute()
	}()
	for i := 1; i < 5; i++ {
		go func() {
			_ = timeSelect.Execute()
		}()
		<-time.After(1 * time.Second)
	}
	go func() {
		_ = timeSelect.Execute()
	}()

	wap := market.NewWAP()
	handler := market.NewHandler()
	handler.WeightedAveragePrice = wap
	req := []market.MarketOperation{
		{Operation: "buy", UnitCost: 20, Quantity: 2},
		{Operation: "buy", UnitCost: 10, Quantity: 9},
	}
	handler.Execute(req)

	go func() {
		_ = timeSelect.Execute()
	}()
	for i := 1; i < 5; i++ {
		go func() {
			_ = timeSelect.Execute()
		}()
		<-time.After(1 * time.Second)
	}
	go func() {
		_ = timeSelect.Execute()
	}()

	<-time.After(5 * time.Second)
}
