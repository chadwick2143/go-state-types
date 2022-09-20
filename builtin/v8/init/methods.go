package init

import (
	"github.com/filecoin-project/go-state-types/abi"
)

var Methods = map[uint64]interface{}{
	1: *new(func(interface{}, *ConstructorParams) *abi.EmptyValue), // Constructor
	2: *new(func(interface{}, *ExecParams) *ExecReturn),            // Exec
}
