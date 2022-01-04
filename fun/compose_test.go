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
	assert.Equal(t, Chain(add5, add5, add5)(10), 25)
	assert.Equal(t, Chain(mul10, add5, add5)(10), 110)
	assert.Equal(t, Chain(add5, mul10, div3)(10), 50)
}

func TestCompose(t *testing.T) {
	assert.Equal(t, Compose(mul10, add5)(10), 105)
	assert.Equal(t, Compose(add5, mul10)(10), 150)
	assert.Equal(t, Compose(atoi, add5)("10"), 15)
}

func TestCompose2(t *testing.T) {
	assert.Equal(t, Compose2(mul10, add5)(10), 105)
	assert.Equal(t, Compose2(add5, mul10)(10), 150)
	assert.Equal(t, Compose2(atoi, add5)("10"), 15)
}

func TestCompose3(t *testing.T) {
	assert.Equal(t, Compose3(fmulpi, ftoi, itoa)(10), "31")
}

func TestCompose4(t *testing.T) {
	assert.Equal(t, Compose4(fmulpi, ftoa, atof, mul7i)(10), 219)
}
