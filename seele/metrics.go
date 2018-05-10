package seele

import (
	"github.com/seeleteam/go-seele/metrics"
)

var (
	blockHeight = metrics.NewRegisteredGauge("seele.blockHeight", nil)
)
