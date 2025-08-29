package glua

import (
	"context"
	lua "github.com/yuin/gopher-lua"
)


// Call func with return
func Call(L *lua.LState, fn lua.LValue, opts ...Option) (lua.LValue, error) {
	var callOptions options
	for _, opt := range opts {
		opt(&callOptions)
	}
	p := lua.P{
		Fn: fn,
		NRet: 1,
		Protect: callOptions.protect,
		Handler: callOptions.handler,
	}
	if callOptions.ctx != nil {
		L.SetContext(callOptions.ctx)
	}
	if err := L.CallByParam(p, callOptions.args...); err != nil {
		return nil, err
	}
	v := L.Get(-1)
	L.Pop(1)
	return v, nil
}


// callOptions
type options struct {
	ctx 	context.Context
	protect bool
	args 	[]lua.LValue
	handler *lua.LFunction
}

type Option func(*options)

func WithContext(ctx context.Context) Option {
	return func(opt *options) {
		opt.ctx = ctx
	}
}

func WithProtect(protect bool) Option {
	return func(opt *options) {
		opt.protect = protect
	}
}

func WithArgs(args ...lua.LValue) Option {
	return func(opt *options) {
		opt.args = args
	}
}

func WithHandler(handler *lua.LFunction) Option {
	return func(opt *options) {
		opt.handler = handler
	}
}