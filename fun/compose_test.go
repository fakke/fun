package fun

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"strconv"
	"testing"
)

var add5 = func(x int) int { return x + 5 }
var mul10 = func(x int) int { return x * 10 }
var div3 = func(x int) int { return x / 3 }
var fmulpi = func(x int) float32 { return float32(x) * math.Pi }
var mul7i = func(x float32) int { return int(x * 7) }
var ftoa = func(x float32) string { return fmt.Sprintf("%f", x) }
var ftoi = func(x float32) int { return int(x) }
var atof = func(s string) float32 {
	x, e := strconv.ParseFloat(s, 32)
	if e != nil {
		panic(e)
	}
	return float32(x)
}
var itoa = strconv.Itoa
var atoi = func(s string) int {
	x, e := strconv.Atoi(s)
	if e != nil {
		panic(e)
	}
	return x
}

func TestChain(t *testing.T) {
	assert.Equal(t, 25, Chain(add5, add5, add5)(10))
	assert.Equal(t, 110, Chain(mul10, add5, add5)(10))
	assert.Equal(t, 50, Chain(add5, mul10, div3)(10))
}

func TestCompose(t *testing.T) {
	assert.Equal(t, 105, Compose(mul10, add5)(10))
	assert.Equal(t, 150, Compose(add5, mul10)(10))
	assert.Equal(t, 15, Compose(atoi, add5)("10"))
}

func TestCompose2(t *testing.T) {
	assert.Equal(t, 105, Compose2(mul10, add5)(10))
	assert.Equal(t, 150, Compose2(add5, mul10)(10))
	assert.Equal(t, 15, Compose2(atoi, add5)("10"))
}

func TestCompose3(t *testing.T) {
	assert.Equal(t, "31", Compose3(fmulpi, ftoi, itoa)(10))
}

func TestCompose4(t *testing.T) {
	assert.Equal(t, 219, Compose4(fmulpi, ftoa, atof, mul7i)(10))
}

func TestComposeAny(t *testing.T) {
	//f1 := func(x float32) int { return 1 }
	//f2 := func(x int) string { return "2" }
	fs := make([]func(any) any, 2, 2)
	assert.Equal(t, len(fs), 2)
	//assert.Equal(t, 219, Compose4(fmulpi, ftoa, atof, mul7i)(10))
}
