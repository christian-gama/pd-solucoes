package bench

import (
	"time"

	"github.com/christian-gama/pd-solucoes/pkg/log"
)

// Duration returns the duration that a function takes to execute.
func Duration(fn func()) time.Duration {
	start := time.Now()
	fn()
	elapsed := time.Since(start)

	return elapsed
}

// PrintDuration prints the duration that a function takes to execute.
func PrintDuration(log log.Logger, resource string, fn func()) {
	duration := Duration(fn)
	log.Infof("%s took %dms to complete", resource, duration.Milliseconds())
}
