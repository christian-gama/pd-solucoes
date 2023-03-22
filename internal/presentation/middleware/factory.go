package middleware

import "github.com/christian-gama/pd-solucoes/pkg/log"

func MakeError() Error {
	return NewError(log.MakeLogWithCaller(1))
}
