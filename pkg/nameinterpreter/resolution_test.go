package nameinterpreter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResolutionRegex(t *testing.T) {
	title := "Dick Turpin S01E01-02 1080p WEB-DL DUAL 5 1"
	assert.Equal(t, "1080p", RESOLUTION_REG.FindString(title))
}
