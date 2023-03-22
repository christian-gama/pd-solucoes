package validators

import (
	"fmt"
	"regexp"

	"github.com/christian-gama/pd-solucoes/internal/infra/querying"
	"github.com/christian-gama/pd-solucoes/pkg/slice"
)

// Filter returns true if the string is a valid filter field.
// For example: "field:name,eq:John" or "field:name,like:John" or "field:name,in:[1 2 3]".
func Filter(v string, params []string) bool {
	reg := fmt.Sprintf(`field=(\w+),(%s)=(\w+|\[[^\]]*\])`, querying.AllowedFilterOperators())

	if !regexp.MustCompile(reg).MatchString(v) {
		return false
	}

	fieldName := regexp.MustCompile(reg).FindStringSubmatch(v)[1]
	if fieldName == "" {
		return false
	}

	if !slice.Contains(params, fieldName) {
		return false
	}

	value := regexp.MustCompile(reg).FindStringSubmatch(v)[3]
	if value == "" {
		return false
	}

	// Array is only allowed for the "in" operator.
	if regexp.MustCompile(`\[.*\]`).MatchString(value) &&
		regexp.MustCompile(reg).FindStringSubmatch(v)[2] != "in" {
		return false
	}

	return true
}
