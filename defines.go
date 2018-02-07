package max31856

import (
	"github.com/mvpninjas/go-bitflag"
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

