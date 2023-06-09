package env_test

import (
	"fmt"
	"os"
	gopath "path"
	"testing"
	"time"

	"github.com/christian-gama/pd-solucoes/internal/infra/env"
	"github.com/christian-gama/pd-solucoes/pkg/path"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
	testify "github.com/stretchr/testify/suite"
)

type LoaderSuite struct {
	suite.Suite
}

func TestLoaderSuite(t *testing.T) {
	testify.Run(t, new(LoaderSuite))
}

func (s *LoaderSuite) TestLoad() {
	s.Run("loads the environment variables", func() {
		s.NotPanics(func() {
			os.Clearenv()
			file, envFileName := createTempEnv(validEnvContent)
			defer os.Remove(file.Name())

			env.Load(envFileName)

			s.NotZero(env.App.Env, "env.App.Env")
			s.NotZero(env.App.Port, "env.App.Port")
			s.NotZero(env.App.Host, "env.App.Host")
			s.NotZero(env.Config.GlobalRateLimit, "env.Config.GlobalRateLimit")
			s.NotZero(env.Config.Debug, "env.Config.Debug")
			s.NotZero(env.DB.Host, "env.DB.Host")
			s.NotZero(env.DB.Name, "env.DB.Name")
			s.NotZero(env.DB.Password, "env.DB.Password")
			s.NotZero(env.DB.Port, "env.DB.Port")
			s.NotZero(env.DB.User, "env.DB.User")
			s.NotZero(env.DB.SslMode, "env.DB.SslMode")
		})

		s.Panics(func() {
			env.Load("invalid")
		})
	})
}

func createTempEnv(content string) (*os.File, string) {
	randomStr := fmt.Sprintf("%d", time.Now().UnixNano())
	envFileName := fmt.Sprintf(".env.temp.%s", randomStr)
	rootDir := path.Root()

	file, err := os.Create(gopath.Join(rootDir, envFileName))
	if err != nil {
		panic(err)
	}

	_, err = file.WriteString(content)
	if err != nil {
		os.Remove(file.Name())
		panic(err)
	}

	return file, envFileName
}

const validEnvContent = `
APP_ENV=test
APP_PORT=123
APP_HOST=localhost
CONFIG_GLOBAL_RATE_LIMIT=10
CONFIG_DEBUG=true
DB_HOST=localhost
DB_NAME=test
DB_PASSWORD=123
DB_PORT=123
DB_SGBD=postgresql
DB_SSL_MODE=disable
DB_USER=123
`
