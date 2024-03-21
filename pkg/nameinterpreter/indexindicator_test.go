package nameinterpreter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexIndicatorRegex(t *testing.T) {
	title := "Dick Turpin S01E01-02 1080p WEB-DL DUAL 5 1"
	matches := INDEX_INDICATOR_REG.FindAllString(title, 2)
	assert.Equal(t, "S01", matches[0])
	assert.Equal(t, "E01-02", matches[1])
}

func TestIndexIndicatorParser(t *testing.T) {
	prefix, indexes := ParseIndexIndicator("E01-02-03")
	assert.Equal(t, "E", prefix)
	assert.Equal(t, 1, indexes[0])
	assert.Equal(t, 2, indexes[1])
	assert.Equal(t, 3, indexes[2])
}

