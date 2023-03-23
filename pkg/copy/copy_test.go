package copy_test

import (
	"errors"
	"testing"

	"github.com/christian-gama/pd-solucoes/pkg/copy"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
)

type CopySuite struct {
	suite.Suite
}

func TestCopy(t *testing.T) {
	suite.RunUnitTest(t, new(CopySuite))
}

func (s *CopySuite) TestCopy() {
	s.Run("copy from source to destination", func() {
		source := &CopySource{Name: "John", Age: 25}
		destination := &CopyDestination{}

		result := copy.MustCopy(destination, source)

		s.Equal(source.Name, result.Name)
		s.Equal(source.Age, result.Age)
	})

	s.Run("panics if destination is not a pointer", func() {
		source := &CopySource{Name: "John", Age: 25}
		destination := CopyDestination{}

		s.PanicsWithError("copy destination is invalid", func() {
			copy.MustCopy(destination, source)
		})
	})

	s.Run("copy even if source is not a pointer", func() {
		source := CopySource{Name: "John", Age: 25}
		destination := &CopyDestination{}

		copy.MustCopy(destination, source)

		s.Equal(source.Name, destination.Name)
		s.Equal(source.Age, destination.Age)
	})
}

type CopySource struct {
	Name string
	Age  int
}

type CopyDestination struct {
	Name string
	Age  int
}

func (d *CopyDestination) Validate() error {
	if d.Name == "" {
		return errors.New("name is required")
	}

	return nil
}
