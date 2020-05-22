package pagination

import (
	"os"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

// MustHaveEnv ensures the ENV exists, otherwise will crash the application
func MustHaveEnv(key string) string {
	env := os.Getenv(key)
	if env == "" {
		logrus.Fatalf("%s variable is not set", key)
	}
	return env
}

// MustHaveEnvInt ensures the ENV exists and returns it as an integer, otherwise will crash the application
func MustHaveEnvInt(key string) int {
	env := MustHaveEnv(key)
	envInt, err := strconv.Atoi(env)
	if err != nil {
		logrus.Fatal(err)
	}
	return envInt
}
