package validators_test

import (
	"testing"

	"github.com/christian-gama/pd-solucoes/internal/infra/validation/validators"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
)

type FilterSuite struct {
	suite.Suite
}

func TestFilterSuite(t *testing.T) {
	suite.RunUnitTest(t, new(FilterSuite))
}

func (s *FilterSuite) TestFilter() {
	s.Run("returns true if the filter is valid", func() {
		s.True(validators.Filter("field=name,eq=John", []string{"name"}))
		s.True(validators.Filter("field=name,like=John", []string{"name"}))
		s.True(validators.Filter("field=name,in=[1 2 3]", []string{"name"}))
		s.True(validators.Filter("field=name,gt=1", []string{"name"}))
		s.True(validators.Filter("field=name,gte=1", []string{"name"}))
		s.True(validators.Filter("field=name,lt=1", []string{"name"}))
		s.True(validators.Filter("field=name,lte=1", []string{"name"}))
		s.True(validators.Filter("field=name,neq=1", []string{"name"}))
		s.True(validators.Filter("field=name,eq=John", []string{"name", "id"}))
		s.True(validators.Filter("field=name,like=John", []string{"name", "id"}))
		s.True(validators.Filter("field=name,in=[1 2 3]", []string{"name", "id"}))
		s.True(validators.Filter("field=name,gt=1", []string{"name", "id"}))
		s.True(validators.Filter("field=name,gte=1", []string{"name", "id"}))
		s.True(validators.Filter("field=name,lt=1", []string{"name", "id"}))
		s.True(validators.Filter("field=name,lte=1", []string{"name", "id"}))
		s.True(validators.Filter("field=name,neq=1", []string{"name", "id"}))
	})

	s.Run("returns false if the filter is invalid", func() {
		s.False(validators.Filter("field=name,eq=John", []string{"id"}))
		s.False(validators.Filter("field=name,like=John", []string{"id"}))
		s.False(validators.Filter("field=name,in=[1 2 3]", []string{"id"}))
		s.False(validators.Filter("field=name,gt=1", []string{"id"}))
		s.False(validators.Filter("field=name,gte=1", []string{"id"}))
		s.False(validators.Filter("field=name,lt=1", []string{"id"}))
		s.False(validators.Filter("field=name,lte=1", []string{"id"}))
		s.False(validators.Filter("field=name,neq=1", []string{"id"}))
		s.False(validators.Filter("field=name,invalid=1", []string{"name"}))
		s.False(validators.Filter("field=name,eq=John", []string{}))
		s.False(validators.Filter("field=name,in=John", []string{}))
	})
}
