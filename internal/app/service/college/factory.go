package service

import (
	"github.com/christian-gama/pd-solucoes/internal/infra/persistence"
)

func MakeCreateCollege() CreateCollege {
	return NewCreateCollege(persistence.MakeCollege())
}
