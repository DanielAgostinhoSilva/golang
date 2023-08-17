package values_objects

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_deve_criar_um_cpf_valido(t *testing.T) {
	cpf, err := NewCpf("488.440.270-70")
	assert.Nil(t, err)
	assert.Equal(t, "48844027070", cpf.value)
	cpf, err = NewCpf("48844027070")
	assert.Nil(t, err)
	assert.Equal(t, "48844027070", cpf.value)
}

func Test_deve_lancar_um_erro_quando_o_cpf_for_invalido(t *testing.T) {
	_, err := NewCpf("")
	assert.Error(t, err)
	_, err = NewCpf("488.440.270")
	assert.Error(t, err)
	_, err = NewCpf("488.440.270-71")
	assert.Error(t, err)
	_, err = NewCpf("488.440.270-10")
	assert.Error(t, err)
}
