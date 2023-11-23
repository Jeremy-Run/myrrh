package business

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEval_LoginTimes(t *testing.T) {
	result := Eval("LoginTimes")
	assert.Equal(t, result, 5)
}
