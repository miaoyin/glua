package glua

import lua "github.com/yuin/gopher-lua"

func PrintReg(L *lua.LState) error {
	return L.DoString(`_printregs()`)
}
