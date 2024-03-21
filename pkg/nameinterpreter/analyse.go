// TODO: at the current state it does multiple loop on the same input,
//       planning at using only one loop, but custom parsing will be needed, no regexes

// an torrent media interpreter, it detects seasons, episodes,
// resolutions, sources and codecs if too ambiguous read readme.md
package nameinterpreter

import "strings"

type Info struct {
   Seasons []int
   Episodes []int
   Resolution string
   Provider string
   Quality string
}

func inspectIIMatches(i *Info, matches []string) {
   for _, match := range matches {
      switch rune(match[0]) {
      case 'S':
         _, indexes := ParseIndexIndicator(string(match[1:]))
         i.Seasons = append(i.Seasons, indexes...)
      case 'E':
         _, indexes := ParseIndexIndicator(string(match[1:]))
         i.Episodes = append(i.Episodes, indexes...)
      }
   }
}

func inspectResolution(i *Info, res string) {
   i.Resolution = res
}

var providers = map[string]string{
   "ATVP": "AppleTV", "DSNP": "Disney+",
}

func identifyProvider(raw string, i *Info) {
   for k, v := range providers {
      if strings.Contains(raw, k) {
         i.Provider = v
         return
      }
   }
}

var qualities = []string{"WEB-DL"}

func identifyQuality(raw string, i *Info) {
   for _, quality := range qualities {
      if strings.Contains(raw, quality) {
         i.Quality = quality
         return
      }
   }
}

func Parse(raw string) *Info {
   info := Info{}
   ii_matches := INDEX_INDICATOR_REG.FindAllString(raw, 2)
   inspectIIMatches(&info, ii_matches)

   resolution := RESOLUTION_REG.FindString(raw)
   inspectResolution(&info, resolution)

   identifyProvider(raw, &info)
   identifyQuality(raw, &info)
   return &info
}

