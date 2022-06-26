package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestLoadConfigurationOrDie_InvalidPath(t *testing.T) {
	defer func() {
		obtained := recover()
		expected := InvalidPath{"../test-resources/doesnt_exist.json", "open ../test-resources/doesnt_exist.json: no such file or directory"}.Error()

		assert.Equal(t, expected, obtained)
	}()

	if err := os.Setenv("CREDENTIALS_PATH", "../test-resources/doesnt_exist.json"); err != nil {
		panic(err)
	}
	LoadConfigurationOrDie()
}

func TestLoadConfigurationOrDie_InvalidField(t *testing.T) {
	defer func() {
		obtained := recover()
		expected := InvalidField{".api.user", "empty string not allowed"}.Error()
		assert.Equal(t, expected, obtained)
	}()

	if err := os.Setenv("CREDENTIALS_PATH", "../test-resources/invalid_field.json"); err != nil {
		panic(err)
	}
	LoadConfigurationOrDie()
}
