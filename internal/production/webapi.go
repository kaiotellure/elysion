package production

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/ikaio/tailmplx/internal/database"
	"github.com/ikaio/tailmplx/internal/help"
	"github.com/muesli/gamut"
)

func markdownToHTML(md []byte) []byte {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}

func HandlePut(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "body reading failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var production database.Production
	err = json.Unmarshal(body, &production)
	if err != nil {
		http.Error(w, "json parsing failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	production.PostProcess.DescriptionHTML = string(markdownToHTML([]byte(production.Description)))

	primary, err := help.GetImagePrimaryColorFromURL(production.Images.Cover)
	if err != nil {
		http.Error(w, "cover image primary color detection failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	production.PostProcess.PrimaryColor = primary

	color := gamut.Hex(primary)
	production.PostProcess.LigherColor = gamut.ToHex(gamut.Lighter(color, 1.))
	production.PostProcess.DarkerColor = gamut.ToHex(gamut.Darker(color, .4))

	err = production.Save()
	if err != nil {
		http.Error(w, "database saving failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Production was saved successfully."))
}
