package erro

import (
	"bytes"
	"context"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"strconv"
)

// Big error
type Big struct {
	Code int              `json:"code,omitempty"`
	Text string           `json:"error,omitempty"`
	Data interface{}      `json:"data,omitempty"`
	Ctx  *context.Context `json:"-"`
}

// NewBig returns a new Big error with data
func NewBig(data interface{}) *Big {
	big := new(Big)

	if data == nil {
		return nil
	}

	if s, ok := data.(string); ok {
		big.Text = s
	}

	big.Data = data
	return big
}

func NewBigCtx(ctx *context.Context, data interface{}) *Big {
	big := NewBig(data)
	big.Ctx = ctx
	return big
}

// Error error
func (b *Big) Error() string {
	if b.Text != "" {
		return b.Text
	}
	if b.Data != nil {
		if s, ok := b.Data.(string); ok {
			return s
		}
		if i, ok := b.Data.(int); ok {
			return strconv.Itoa(i)
		}
		if i, ok := b.Data.(int64); ok {

			return strconv.FormatInt(i, 10)
		}
		if f, ok := b.Data.(float64); ok {
			return strconv.FormatFloat(f, 'f', -1, 32)
		}
	}

	return ""
}

// MarshalJSON bytes
func (b *Big) MarshalJSON() []byte {
	j, _ := json.MarshalIndent(b, " ", " ")
	return j
}

// MarshalBinary bytes
func (b *Big) MarshalBinary() []byte {
	var buf = new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	enc.Encode(b.Data)
	databytes := buf.Bytes()
	buf.Reset()
	fmt.Fprintf(buf, "%b\x00\x00\x00%b", []byte(b.Text), databytes)
	return buf.Bytes()
}
