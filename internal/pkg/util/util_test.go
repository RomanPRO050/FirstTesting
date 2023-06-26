package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseInt(t *testing.T) {
	number, err := ReverseInt(123)
	assert.Nil(t, err)
	assert.Equal(t, number, 321, "Число должно быть наоборот")
	NegNumber, err1 := ReverseInt(-649)
	assert.Nil(t, err1)
	assert.Equal(t, NegNumber, -946, "Число должно быть наоборот")
	ZeroNumber, err2 := ReverseInt(0)
	assert.Nil(t, err2)
	assert.Equal(t, ZeroNumber, 0, "Число должно быть равно 0")
	_, err3 := ReverseInt(2147483648)
	assert.Nil(t, err3, "За пределами значений")
	_, err4 := ReverseInt(2.1)
	assert.NotNil(t, err4, "число не INT")
}
