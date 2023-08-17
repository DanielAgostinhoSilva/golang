package values_objects

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_deve_criar_um_nome_valido(t *testing.T) {
	name, err := NewName("Test A")
	assert.Nil(t, err)
	assert.Equal(t, "Test A", name.GetValue())
}

func Test_deve_verificar_se_um_Name_e_igual_ao_outro(t *testing.T) {
	nameA, _ := NewName("Test A")
	nameB, _ := NewName("Test A")
	assert.True(t, nameA.Equals(*nameB))
}

func Test_deve_lancar_um_erro_quando_nome_estiver_em_branco(t *testing.T) {
	_, err := NewName("")
	assert.Equal(t, ErrNameIsRequired, err)
}

func Test_deve_lancar_um_erro_quando_nome_estiver_invalido(t *testing.T) {
	_, err := NewName("ab")
	assert.Equal(t, ErrInvalidName, err)
}
