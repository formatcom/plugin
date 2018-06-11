package main

import (
	"os"
	"fmt"
	"unsafe"
	"github.com/formatcom/plugin"
)

type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

func main(){
	plugin.Test()
	os.Exit(0)


	fmt.Println("hola")
	p := plugin.Plugin{}
	fmt.Println("hola 2")
	if err := p.Open("plugin/plugin.so"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("hola 3")
	sym, err := p.Lookup("plugin_func");
	fmt.Println("hola 4")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%p\n", &sym)
	fmt.Printf("%p\n", unsafe.Pointer(&sym))
	f := *(*func())( unsafe.Pointer(&sym) )
	fmt.Printf("%p\n", &f)
	f()

	fmt.Println(f)

	fmt.Println(p)
}
