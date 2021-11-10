package bizerror

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testCoder struct {
	code int
	msg  string
}

func (tc testCoder) Error() string {
	return tc.msg
}

func (tc testCoder) Message() string {
	return tc.msg
}

func (tc testCoder) Code() int {
	return tc.code
}

func TestMustRegister(t *testing.T) {
	defer func() {
		e := recover()
		assert.True(t, e != nil)
	}()
	e := testCoder{code: 1, msg: "err code 1"}
	MustRegister(1, e)
	MustRegister(1, e)
}

func TestRegister(t *testing.T) {
	errMsg := "err code 1"
	e := testCoder{
		code: 1,
		msg:  errMsg,
	}
	Register(1, e)

	assert.True(t, ParseCoder(e).Message() == errMsg)
}

func TestStack(t *testing.T) {
	err := New(1, "msg")
	bizError := ParseBizError(err)
	t.Log(bizError.Stack())

	bizError2 := ParseBizError(errors.New("msg2"))
	t.Log(bizError2.Code())
}

func TestCoderAndBizError(t *testing.T) {
	errMsg := "err code 1"
	bize := New(1, errMsg)
	codere := testCoder{
		code: 1,
		msg:  errMsg,
	}
	MustRegister(1, codere)
	assert.True(t, ParseBizError(bize).Message() == ParseCoder(codere).Message())
}
