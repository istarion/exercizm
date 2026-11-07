package parsinglogfiles

import "regexp"
import "fmt"

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^(\[TRC\]|\[DBG\]|\[INF\]|\[WRN\]|\[ERR\]|\[FTL\])`)
    return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=-]*>`)
    return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)".*password.*"`)
    sum := 0
    for _, l := range lines {
        if re.MatchString(l) {
            sum++
        }
    }
    return sum
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
    return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\w+)`)
    for i, l := range lines {
        matches := re.FindStringSubmatch(l)
        if matches != nil {
            lines[i] = fmt.Sprintf("[USR] %s %s", matches[1], lines[i])
        }
    }
    return lines
}
