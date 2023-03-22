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
		s.True(validators.Filter("field=name,operator=eq,value=John", []string{"name"}))
		s.True(validators.Filter("field=name,operator=like,value=John", []string{"name"}))
		s.True(validators.Filter("field=name,operator=in,value=[1 2 3]", []string{"name"}))
		s.True(validators.Filter("field=name,operator=gt,value=1", []string{"name"}))
		s.True(validators.Filter("field=name,operator=gte,value=1", []string{"name"}))
		s.True(validators.Filter("field=name,operator=lt,value=1", []string{"name"}))
		s.True(validators.Filter("field=name,operator=lte,value=1", []string{"name"}))
		s.True(validators.Filter("field=name,operator=neq,value=1", []string{"name"}))
		s.True(validators.Filter("field=name,operator=eq,value=John", []string{"name", "id"}))
		s.True(validators.Filter("field=name,operator=like,value=John", []string{"name", "id"}))
		s.True(validators.Filter("field=name,operator=in,value=[1 2 3]", []string{"name", "id"}))
		s.True(validators.Filter("field=name,operator=gt,value=1", []string{"name", "id"}))
		s.True(validators.Filter("field=name,operator=gte,value=1", []string{"name", "id"}))
		s.True(validators.Filter("field=name,operator=lt,value=1", []string{"name", "id"}))
		s.True(validators.Filter("field=name,operator=lte,value=1", []string{"name", "id"}))
		s.True(validators.Filter("field=name,operator=neq,value=1", []string{"name", "id"}))
	})

	s.Run("returns false if the filter is invalid", func() {
		s.False(validators.Filter("field=name,operator=eq,value=John", []string{"id"}))
		s.False(validators.Filter("field=name,operator=like,value=John", []string{"id"}))
		s.False(validators.Filter("field=name,operator=in,value=[1 2 3]", []string{"id"}))
		s.False(validators.Filter("field=name,operator=gt,value=1", []string{"id"}))
		s.False(validators.Filter("field=name,operator=gte,value=1", []string{"id"}))
		s.False(validators.Filter("field=name,operator=lt,value=1", []string{"id"}))
		s.False(validators.Filter("field=name,operator=lte,value=1", []string{"id"}))
		s.False(validators.Filter("field=name,operator=neq,value=1", []string{"id"}))
		s.False(validators.Filter("field=name,invalid=1", []string{"name"}))
		s.False(validators.Filter("field=name,operator=eq,value=John", []string{}))
		s.False(validators.Filter("field=name,operator=in,value=John", []string{}))
	})
}
