package nameinterpreter

import "regexp"

// will match resolutions from 3 to 4 digits followed by a p
// ...720p... WILL NOT match, ... 720p ... WILL match
var RESOLUTION_REG = regexp.MustCompile(`\d{3,4}p`)
