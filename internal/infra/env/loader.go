package env

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/christian-gama/pd-solucoes/pkg/path"
	"github.com/christian-gama/pd-solucoes/pkg/slice"
	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

// Load loads the environment variables from the .env file.
func Load(envFile string) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := godotenv.Load(Path(envFile))
	if err != nil {
		panic(fmt.Errorf("Error loading .env file: %w", err))
	}

	if err := envconfig.Process(ctx, DB); err != nil {
		panic(fmt.Errorf("Error loading DB environment variables: %w", err))
	}

	if err := envconfig.Process(ctx, App); err != nil {
		panic(fmt.Errorf("Error loading App environment variables: %w", err))
	}

	if err := envconfig.Process(ctx, Config); err != nil {
		panic(fmt.Errorf("Error loading Config environment variables: %w", err))
	}

	validate()
}

func validate() {
	validEnvs := []string{Development, Test, Production}
	if !slice.Contains(validEnvs, App.Env) {
		panic(
			fmt.Errorf(
				"Invalid env variable: '%s'. Must be either: %v",
				App.Env,
				validEnvs,
			),
		)
	}

	validSslModes := []string{"disable", "allow", "prefer", "require", "verify-ca", "verify-full"}
	if !slice.Contains(validSslModes, DB.SslMode) {
		panic(
			fmt.Errorf(
				"Invalid env variable: '%s'. Must be either: %v",
				DB.SslMode,
				validSslModes,
			),
		)
	}

	if App.Env == Production {
		if len(DB.Password) < 16 {
			panic(
				fmt.Errorf(
					"Invalid env variable: '%s'. Password must be at least 16 characters long",
					DB.Password,
				),
			)
		}

		if len(DB.User) < 16 {
			panic(
				fmt.Errorf(
					"Invalid env variable: '%s'. User must be at least 16 characters long",
					DB.User,
				),
			)
		}
	}
}

// Path returns the absolute path of the given environment file (envFile) in the Go module's
// root directory. It searches for the 'go.mod' file from the current working directory upwards
// and appends the envFile to the directory containing 'go.mod'.
// It panics if it fails to find the 'go.mod' file.
func Path(envFile string) string {
	rootDir := path.Root()
	return filepath.Join(rootDir, envFile)
}
