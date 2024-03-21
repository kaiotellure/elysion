package production

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/ikaio/tailmplx/help"
	"github.com/muesli/gamut"
)

func MarkdownToHTML(md string) string {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse([]byte(md))

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return string(markdown.Render(doc, renderer))
}

func HandlePut(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "body reading failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var production Production
	err = json.Unmarshal(body, &production)
	if err != nil {
		http.Error(w, "json parsing failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	DoPostProcess(&production)

	err = production.Save()
	if err != nil {
		http.Error(w, "database saving failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Production was saved successfully."))
}

func DoPostProcess(p *Production) error {
	primary, err := help.GetImagePrimaryColorFromURL(p.Images.Cover)
	if err != nil {
		return err
	}

	p.PostProcess.PrimaryColor = primary

	color := gamut.Hex(primary)
	p.PostProcess.LigherColor = gamut.ToHex(gamut.Lighter(color, 1.))
	p.PostProcess.DarkerColor = gamut.ToHex(gamut.Darker(color, .4))
	return nil
}
