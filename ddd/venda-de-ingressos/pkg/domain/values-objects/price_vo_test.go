package values_objects

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_deve_inicializar_um_price_valido(t *testing.T) {
	price, err := NewPrice(10.00)
	assert.Nil(t, err)
	assert.Equal(t, 10.00, price.GetValue())
}

func Test_deve_retornar_um_erro_quando_price_for_menor_que_zero(t *testing.T) {
	price, err := NewPrice(-10.00)
	assert.Nil(t, price)
	assert.Error(t, err)
}
