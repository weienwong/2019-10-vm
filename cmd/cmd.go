package main

import (
	"fmt"
	"io/ioutil"

	"github.com/weienwong/2019-10-vm/virtual"
)

var vm virtual.Machine

func main() {
	f, err := ioutil.ReadFile("../examples/hello.out")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%x\n", f)

	vm.Load(f)

	fmt.Printf("%#v\n", vm)

	if err := vm.Execute(); err != nil {
		fmt.Println(err)
	}
}
