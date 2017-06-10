package erro

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

func TestMarshal(t *testing.T) {
	big := NewBig("im a big error")
	big.Data = fmt.Sprintf("% x", strconv.FormatFloat(math.Pi, 'f', -1, 64))
	b := big.MarshalBinary()
	fmt.Println(string(b))
	b = big.MarshalJSON()
	fmt.Println(string(b))
}
