package bizerror

import (
	"fmt"
	"github.com/gogf/gf/errors/gcode"
	"github.com/gogf/gf/errors/gerror"
	"sync"

	"github.com/pkg/errors"
)

var ErrBiz = errors.New("biz")
var (
	errUnknownMsg  = "Unknown Error"
	errUnknownCode = 1
	errUnknown     = gerror.NewCode(gcode.New(errUnknownCode, errUnknownMsg, ""), "")
	unknownCoder   = BizError{
		C:     errUnknownCode,
		Msg:   errUnknownMsg,
		error: gerror.New(errUnknownMsg),
	}
)
var codeMap = map[int]ApiCoder{}
var codeMux = &sync.Mutex{}

// ApiCoder return to gRPC/HTTP Client
type ApiCoder interface {
	Code() int
	Message() string
	Error() string
}

// BizError 代表业务上捕捉到的错误/异常，可打印堆栈
type BizError struct {
	C     int    // ErrorResponse.ErrCode, return to gRPC/HTTP Client
	Msg   string // ErrorResponse.ErrMsg,  return to gRPC/HTTP Client
	error error  // Underlying error, from 3rd API/Library or errors.Wrap/New/WithMessage
}

func (e BizError) Code() int {
	return e.C
}

func (e BizError) Message() string {
	return e.Msg
}

func (e BizError) Error() string {
	return e.Msg
}

func (e BizError) Stack() string {
	return fmt.Sprintf("%+v", e.error)
}

func (e BizError) Is(target error) bool {
	return target == ErrBiz
}

var NewError = Newf

// ParseBizError 转 ApiCoder
// 如果失败，返回默认错误
func ParseBizError(err error) BizError {
	if v, ok := err.(BizError); ok {
		return v
	}
	return unknownCoder
}

// ParseCoder 从 codeMap 读取 ApiCoder
// 如果未读取到，返回默认错误
func ParseCoder(err error) ApiCoder {
	if err == nil {
		return nil
	}
	if v, ok := err.(ApiCoder); ok {
		if coder, ok := codeMap[v.Code()]; ok {
			return coder
		}
	}
	return unknownCoder
}

// Stack 返回堆栈
// 不存在返回空字符串
func Stack(err error) string {
	if v, ok := err.(BizError); ok {
		return v.Stack()
	}
	return ""
}

// New construct BizError with code and msg
func New(code int, msg string) error {
	return BizError{C: code, Msg: msg, error: gerror.New(msg)}
}

// Newf construct BizError with code, msg and extra message
func Newf(code int, msg string, format string, args ...interface{}) error {
	return BizError{C: code, Msg: msg, error: gerror.Newf(format, args...)}
}

// Wrap construct BizError with code, msg and underlying error
func Wrap(code int, msg string, err error) error {
	return BizError{
		C:     code,
		Msg:   msg,
		error: gerror.Wrap(err, msg),
	}
}

// Wrapf construct BizError with code, msg, underlying error and extra message
func Wrapf(code int, msg string, err error, format string, args ...interface{}) error {
	return BizError{C: code, Msg: msg, error: gerror.Wrapf(err, format, args...)}
}

// Register 注册用户自定义错误码对应的错误
func Register(code int, err ApiCoder) {
	codeMux.Lock()
	defer codeMux.Unlock()
	codeMap[code] = err
}

// MustRegister 注册用户自定义错误码对应的错误
// 如果错误已经存在则 panic
func MustRegister(code int, err ApiCoder) {
	codeMux.Lock()
	defer codeMux.Unlock()
	if _, ok := codeMap[code]; ok {
		panic(fmt.Sprintf("code: %d already exist", code))
	}
	codeMap[code] = err
}

// WithCode return BizError by code
// Example:
//		WithCode(1010000)
func WithCode(code int) error {
	err := errByCode(code)
	return Wrap(code, err.Error(), err)
}

// WithCodef return BizError by code and extra message
// Example:
//		WithCodef(1010000, "extra message: %s", "other")
//		WithCodef(1010000, "extra message")
func WithCodef(code int, format string, args ...interface{}) error {
	err := errByCode(code)
	return Wrap(code, err.Error(), errors.Wrapf(err, format, args...))
}

// WithCodeWrap return BizError by code and error
//
// Example:
//     WithCodeWrap(1001001, err)
func WithCodeWrap(code int, underlyingErr error) error {
	err := errByCode(code)

	return Wrap(code, err.Error(), mergeErr(err, underlyingErr))
}

// WithCodeWrapf return BizError by code, error and extra message
//
// Example:
//     WithCodeWrapf(1001001, err, "关于 err 的更多信息")
//     WithCodeWrapf(1001001, err, "关于 err 的更多信息: %s", "及参数")
func WithCodeWrapf(code int, underlyingErr error, format string, args ...interface{}) error {
	err := errByCode(code)

	return Wrapf(code, err.Error(), mergeErr(err, underlyingErr), format, args...)
}

// Customized return BizError by code and customized message
func Customized(code int, msg string) error {
	err := errByCode(code)
	errMsg := err.Error()
	if msg != "" {
		errMsg = msg
	}
	return Wrap(code, errMsg, err) // Use err as underlying error
}

func errByCode(code int) error {
	err, ok := codeMap[code]
	if !ok {
		return errUnknown
	}
	return err
}

// mergeErr 将 error 上的属性转移到 underlying error 上
func mergeErr(err, underlyingErr error) (retErr error) {
	retErr = underlyingErr

	var decorators []func(error) error
	if IsSilence(err) {
		decorators = append(decorators, Silence)
	}

	for _, decorator := range decorators {
		retErr = decorator(retErr)
	}

	return retErr
}
