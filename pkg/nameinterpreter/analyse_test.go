package nameinterpreter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicAnalysis(t *testing.T) {
   info := Parse("Dick Turpin S01E01-02 WEB-DL 720p DUAL ATVP DDP5.1")
   assert.Contains(t, info.Seasons, 1)
   assert.Contains(t, info.Episodes, 1)
   assert.Contains(t, info.Episodes, 2)
   assert.Equal(t, "720p", info.Resolution)
   assert.Equal(t, "AppleTV", info.Provider)
   assert.Equal(t, "WEB-DL", info.Quality)
}
