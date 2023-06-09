package gintest

import (
	"encoding/json"

	"github.com/christian-gama/pd-solucoes/internal/infra/http"
)

func GetBody(body string) *http.ResponseBody {
	var b *http.ResponseBody
	json.Unmarshal([]byte(body), &b)
	return b
}
