package main

import (
	"os"
	"reflect"
	"testing"
)

func setEnvs(envs map[string]string) (closer func()) {
	originalEnvs := map[string]string{}

	for name, value := range envs {
		if originalValue, ok := os.LookupEnv(name); ok {
			originalEnvs[name] = originalValue
		}
		_ = os.Setenv(name, value)
	}

	return func() {
		for name := range envs {
			_ = os.Unsetenv(name)
		}
	}
}

func TestInitVars(t *testing.T) {
	t.Run("All env vars set correctly", func(t *testing.T) {
		curEnvs := setEnvs(map[string]string{
			"PORT":         "6969",
			"DATABASE_URL": "testing url",
		})
		defer curEnvs()

		expected := map[string]string{
			"port":  "6969",
			"dbUrl": "testing url",
		}

		result, _ := initVars()
		if reflect.DeepEqual(result, expected) == false {
			t.Fatalf("Expected %v, got %v\n", expected, result)
		}
	})

	t.Run("One env vars set correctly", func(t *testing.T) {
		curEnvs := setEnvs(map[string]string{
			"PORT": "6969",
		})
		defer curEnvs()

		result, _ := initVars()
		if result != nil {
			t.Fatalf("Expected %v, got %v\n", nil, result)
		}
	})

	t.Run("No env vars set correctly", func(t *testing.T) {
		result, _ := initVars()
		if result != nil {
			t.Fatalf("Expected %v, got %v\n", nil, result)
		}
	})
}
