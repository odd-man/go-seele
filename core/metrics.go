package core

import (
	"github.com/seeleteam/go-seele/metrics"
)

var (
	chainBlockInsert = metrics.NewRegisteredMeter("chain.block.insert", nil)
)
