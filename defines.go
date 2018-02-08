package max31856

import (
	"github.com/mvpninjas/go-bitflag"
)

// Read-only register names
const (
	CR0_RD = iota
	CR1_RD
	MASK_RD
	CJHF_RD
	CJLF_RD
	LTHFTH_RD
	LTHFTL_RD
	LTLFTH_RD
	LTLFTL_RD
	CJTO_RD
	CJTH_RD
	CJTL_RD
	LTCBH_RD
	LTCBM_RD
	LTCBL_RD
	SR_RD
)

// Write-only register names
const (
	CR0_WR = 0x80 + iota
	CR1_WR
	MASK_WR
	CJHF_WR
	CJLF_WR
	LTHFTH_WR
	LTHFTL_WR
	LTLFTH_WR
	LTLFTL_WR
	CJTO_WR
	CJTH_WR
	CJTL_WR
)

// Bit files for register 0x00: Configuration 0 Register
const (
	NR50_60 bitflag.Flag = 1 << bitflag.Flag(iota)
	FAULTCLR
	FAULT
	CJ
	OCFAULT0
	OCFAULT1
	ONESHOT
	CMODE
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
