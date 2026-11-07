package logs

import "strings"
import "unicode/utf8"

const RecommendationMark = rune(0x2757)
const SearchMark = rune(0x1F50D)
const WeatherMark = rune(0x2600)

func Application(log string) string {
    for _, c := range log {
        switch c {
        case RecommendationMark: return "recommendation"
        case SearchMark: return "search"
        case WeatherMark: return "weather"
        }
    }
    return "default"
}

func Replace(log string, oldRune, newRune rune) string {
	return strings.ReplaceAll(log, string(oldRune), string(newRune))
}

func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
