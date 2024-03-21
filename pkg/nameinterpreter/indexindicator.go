package nameinterpreter

import (
	"regexp"
	"strconv"
	"strings"
)

// will match: S01E01-02 --> [S01][E01-02]
var INDEX_INDICATOR_REG = regexp.MustCompile(`[S|E][\d-]+`)

// given E01-02-03, returns prefix=E indexes=[1,2,3]
func ParseIndexIndicator(raw string) (prefix string, indexes []int) {
	prefix = string(raw[0])
	for _, index := range strings.Split(raw[1:], "-") {
		number, _ := strconv.Atoi(index)
		indexes = append(indexes, number)
	}
	return
}
