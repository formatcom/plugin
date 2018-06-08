// ref: https://golang.org/src/plugin/
package "plugin"

import (
	"fmt"
	"errors"
)

import "C"

type Symbol interface{}
type Plugin struct {
	handle C.uintptr_t
}

func (p *Plugin) Open(path string) (*Plugin, error){
	var cErr *C.char
	handle := C.pluginOpen(C.CString(path), &cErr);
	if handle == 0 {
		return nil, errors.New(`plugin.Open("` + path + `"): ` + C.GoString(cErr))
	}
	return Plugin{handle}, nil
}

func (p *Plugin) Lookup(symName string) (Symbol, error){
	var cErr *C.char
	sym := C.pluginLookup(p.handle, C.CString(symName), &cErr)
	if sym == nil {
		return nil, errors.New(`plugin.Open("` + name + `"): could not find symbol ` + symName + `: ` + C.GoString(cErr))
	}
	return sym, nil
}
