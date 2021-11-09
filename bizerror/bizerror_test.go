package bizerror

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMustRegister(t *testing.T) {
	defer func() {
		e := recover()
		assert.True(t, e != nil)
	}()
	MustRegister(1, errors.New("err code 1"))
	MustRegister(2, errors.New("err code 2"))
	MustRegister(1, errors.New("err code 1"))
}

func TestRegister(t *testing.T) {
	errMsg := "err code 1"
	Register(1, errors.New(errMsg))
	codeErrMsg := WithCode(1).(*BizError).Message()
	assert.True(t, New(1, errMsg).(*BizError).Msg == codeErrMsg)
}

func TestNew(t *testing.T) {
	errMsg := "err code 1"
	code := 1
	err := New(code, errMsg)
	assert.True(t, err.Error() == fmt.Sprintf(errFormat, code, errMsg))
}
