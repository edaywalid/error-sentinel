package utils

import (
	"fmt"
	"path/filepath"
	"strings"
)

func FormatStackTrace(stack string) string {
	lines := strings.Split(stack, "\n")
	var formatted strings.Builder
	for _, line := range lines {
		if strings.HasPrefix(line, "\t") {
			parts := strings.Split(line, ":")
			if len(parts) >= 2 {
				file := parts[0][1:]
				lineNumber := strings.Split(parts[1], " ")[0]
				formatted.WriteString(fmt.Sprintf(`<a href="/source?file=%s&line=%s">%s:%s</a><br>`, file, lineNumber, filepath.Base(file), lineNumber))
			}
		} else {
			formatted.WriteString(line + "<br>")
		}
	}
	return formatted.String()
}
