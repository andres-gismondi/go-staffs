package pkg

import (
	"encoding/json"
	"fmt"
	"io"
)

type Decode struct {
	Decoder *json.Decoder
}

func (d Decode) Decode(v interface{}) error {
	// deca := json.NewDecoder(os.Stdin)
	err := d.Decoder.Decode(v)
	if err != nil {
		return fmt.Errorf("error decoding interface: %w", err)
	}
	return nil
}

func (d Decode) WriteRequest(request interface{}, f io.StringWriter) {
	req, err := json.Marshal(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	reqJSON := fmt.Sprintf("Operations: %s \n", req)
	_, _ = f.WriteString(reqJSON)
}

func (d Decode) WriteResponse(res interface{}, f io.StringWriter) {
	b, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
		return
	}

	result := fmt.Sprintf("Taxes: %s \n", b)
	_, _ = f.WriteString(result)
}
