package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

func SourceHandler(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Query().Get("file")
	lineStr := r.URL.Query().Get("line")

	line, err := strconv.Atoi(lineStr)
	if err != nil {
		log.Error().Msg("Invalid line number")
		http.Error(w, "Invalid line number", http.StatusBadRequest)
		return
	}

	content, err := os.ReadFile(file)
	if err != nil {
		log.Error().Msgf("Unable to read file %s: ", file)
		http.Error(w, "Unable to read file: "+file, http.StatusInternalServerError)
		return
	}

	renderSource(w, r, string(content), line)
}

func renderSource(w http.ResponseWriter, r *http.Request, content string, highlightLine int) {
	w.Header().Set("Content-Type", "text/html")
	lines := strings.Split(content, "\n")

	fmt.Fprintf(w, "<h1>Viewing file: %s</h1>", filepath.Base(r.URL.Query().Get("file")))
	fmt.Fprintf(w, "<pre>")

	for i, line := range lines {
		if i+1 == highlightLine {
			fmt.Fprintf(w, `<span id="highlighted-line" style="background-color: yellow">%5d: %s</span><br>`, i+1, template.HTMLEscapeString(line))
		} else {
			fmt.Fprintf(w, "%5d: %s<br>", i+1, template.HTMLEscapeString(line))
		}
	}

	fmt.Fprintf(w, "</pre>")
	fmt.Fprintf(w, `
        <script>
            window.onload = function() {
                var element = document.getElementById("highlighted-line");
                if (element) {
                    element.scrollIntoView({ behavior: "smooth", block: "center" });
                }
            }
        </script>
    `)
}
