package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/takaaki-mizuno/go-boilerplate/config"
)

func TestNewLogger(t *testing.T) {
	configInstance, _ := config.LoadTestConfig()
	t.Run("create logger with test config", func(t *testing.T) {
		logger, err := NewLogger(configInstance)
		assert.NoError(t, err, "return no error")
		assert.NotNil(t, logger, "logger is not null")
	})
}
