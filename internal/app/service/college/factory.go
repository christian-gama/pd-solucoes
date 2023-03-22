package service

import (
	"github.com/christian-gama/pd-solucoes/internal/infra/persistence"
)

func MakeCreateCollege() CreateCollege {
	return NewCreateCollege(persistence.MakeCollege())
}

func MakeUpdateCollege() UpdateCollege {
	return NewUpdateCollege(persistence.MakeCollege())
}

func MakeFindOneCollege() FindOneCollege {
	return NewFindOneCollege(persistence.MakeCollege())
}

func MakeFindAllColleges() FindAllColleges {
	return NewFindAllColleges(persistence.MakeCollege())
}
