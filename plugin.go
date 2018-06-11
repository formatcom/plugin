// ref: https://golang.org/src/plugin/
package plugin

import (
	"fmt"
	"errors"
)

import "C"

type Symbol interface{}
type Plugin struct {
	name string
	handle C.uintptr_t
	Handle C.uintptr_t
}

func Test() {
	p := C.test2()
	fmt.Println(unsafe.Pointer(&p))
}

func (p *Plugin) Open(path string) error {
	p.name = path
	var cErr *C.char
	handle := C.pluginOpen(C.CString(path), &cErr);
	if handle == 0 {
		return errors.New(`plugin.Open("` + p.name + `"): ` + C.GoString(cErr))
	}
	p.handle = handle
	p.Handle = handle
	return nil
}

func (p *Plugin) Lookup(symName string) (Symbol, error){
	var cErr *C.char
	sym := C.pluginLookup(p.handle, C.CString(symName), &cErr)
	if sym == nil {
		return nil, errors.New(`plugin.Open("` + p.name + `"): could not find symbol ` + symName + `: ` + C.GoString(cErr))
	}
	return sym, nil
}
