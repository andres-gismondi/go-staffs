package market

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
)

const (
	taxLimit = 20000
	taxRate  = 0.2
)

var losses = float64(0)

type Handler struct {
	Decoder interface {
		Decode(v interface{}) error
	}
	RequestWriter interface {
		WriteRequest(request interface{}, f io.StringWriter)
	}
	ResponseWriter interface {
		WriteResponse(res interface{}, f io.StringWriter)
	}
	WeightedAveragePrice *WeightedAveragePrice
	operation            map[string]func(buy MarketOperation) TaxResponse
}

func NewHandler() *Handler {
	handler := &Handler{}
	mapOperation := make(map[string]func(MarketOperation) TaxResponse)
	mapOperation["buy"] = handler.Buy
	mapOperation["sell"] = handler.Sell
	handler.operation = mapOperation
	return handler
}

func (h *Handler) Execute(operations []MarketOperation) {
	fmt.Println("Provide operations...")
	f, err := os.Create("data.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			return
		}
	}(f)

	//for {
	//var operations []MarketOperation
	/*err := h.Decoder.Decode(&operations)
	if err != nil {
		log.Fatal(err)
		return
	}*/

	//h.RequestWriter.WriteRequest(&operations, f) //nolint:typecheck

	var taxes []TaxResponse //nolint:prealloc
	for _, operation := range operations {
		taxes = append(taxes, h.operation[operation.Operation](operation))
	}

	//h.ResponseWriter.WriteResponse(&taxes, f) //nolint:typecheck
	//}
}

func (h *Handler) Buy(buy MarketOperation) TaxResponse {
	h.WeightedAveragePrice.addPurchase(buy.UnitCost, buy.Quantity)
	return TaxResponse{Tax: 0}
}

func (h *Handler) Sell(sell MarketOperation) TaxResponse {
	var tax float64

	sellValue := sell.UnitCost * float64(sell.Quantity)
	costValue := h.WeightedAveragePrice.value * float64(sell.Quantity)

	tax = h.setTaxes(sellValue, costValue)
	if tax == 0 {
		tax = h.setLosses(sellValue, costValue)
	}
	h.WeightedAveragePrice.updateQuantity(sell.Quantity)
	return TaxResponse{Tax: tax}
}

func (h *Handler) setTaxes(sellValue float64, costValue float64) float64 {
	if sellValue < costValue {
		return 0
	}

	if sellValue > taxLimit && losses < 1 {
		return (sellValue - costValue) * taxRate
	}

	if sellValue > taxLimit && losses > 0 {
		difference := (sellValue - costValue) - losses
		h.resetLosses(difference)
		if difference < 0 {
			return 0
		}
		return difference * taxRate
	}
	return 0
}

func (h *Handler) setLosses(sellValue float64, costValue float64) float64 {
	if sellValue < costValue {
		losses = costValue - sellValue
	}

	if sellValue > costValue && sellValue < taxLimit {
		total := sellValue - costValue
		h.resetLosses(total - losses)
	}
	return 0
}

func (h *Handler) resetLosses(difference float64) {
	if difference < 0 {
		losses = math.Abs(difference)
	} else {
		// no more debs
		losses = 0
	}
}
