package model_security

import "math"

type ReadAccessLevelsEnum uint32

const (
	None          ReadAccessLevelsEnum = 0
	GenericFields ReadAccessLevelsEnum = 1 << iota
	All           ReadAccessLevelsEnum = (1 << 3) - 1
)

const (
	FULL_DENY              = uint32(math.MaxUint32) // uint32 var with all bits set to 1
	DENY_ACCESS_LEVEL_READ = FULL_DENY - uint32(GenericFields)
)
