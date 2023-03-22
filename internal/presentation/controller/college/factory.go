package controller

import service "github.com/christian-gama/pd-solucoes/internal/app/service/college"

func MakeCreateCollege() CreateCollege {
	return NewCreateCollege(service.MakeCreateCollege())
}

func MakeUpdateCollege() UpdateCollege {
	return NewUpdateCollege(service.MakeUpdateCollege())
}

func MakeFindOneCollege() FindOneCollege {
	return NewFindOneCollege(service.MakeFindOneCollege())
}

func MakeFindAllColleges() FindAllColleges {
	return NewFindAllColleges(service.MakeFindAllColleges())
}
