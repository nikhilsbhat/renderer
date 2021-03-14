package buildercli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getEnvs(t *testing.T) {
	t.Run("should return converted map of environment varaiables", func(t *testing.T) {
		expected := map[string]string{
			"BUILD_ENVIRONMENT": "production",
			"GOVERSION":         "go1.15.7",
		}

		actual, err := getEnvs([]string{"BUILD_ENVIRONMENT=production", "GOVERSION=go1.15.7"})

		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})
}
