package main

import (
	"fmt"
	"github.com/mvpninjas/go-bitflag"
	"reflect"
)

// Bit fields for register 0x0F: Fault Status Register
const (
	OPEN bitflag.Flag = 1 << bitflag.Flag(iota)
	OVUV
	TCLOW
	TCHIGH
	CJLOW
	CJHIGH
	TCRange
	CJRange
)

func main() {
	var flag bitflag.Flag

	flag.Set(TCHIGH)
	fmt.Println(flag)
	fmt.Println(reflect.TypeOf(flag))
	fmt.Println(byte(flag))

	flag.Set(OPEN)
	fmt.Println(flag)
}
