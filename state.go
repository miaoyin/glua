package glua

import (
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

func SetGlobal(L *lua.LState, name string, value any) {
	L.SetGlobal(name, luar.New(L, value))
}
