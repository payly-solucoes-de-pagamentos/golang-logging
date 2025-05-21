package logging

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLogger_WhenCalled_ShouldCreateInstanceOfLogger(t *testing.T) {
	// arrange
	var l *Logger

	// act
	func() {
		l = NewLogger()
	}()

	// assert
	assert.IsType(t, Log, l)
}
