package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/ikaio/tailmplx/components"
	"github.com/ikaio/tailmplx/help"
	"github.com/ikaio/tailmplx/services/crawler"
	"github.com/ikaio/tailmplx/services/database"
	"github.com/ikaio/tailmplx/services/production"
)

const description_fmt = "# Sinopse\n%s\n\n# Informações\n%s"

func metadataToMarkdownList(m map[string]string) string {
	b := strings.Builder{}
	for k, v := range m {
		b.WriteString(fmt.Sprintf("- `%s` %s\n", k, v))
	}
	return b.String()
}

func convertComandoProductionToProduction(cp *crawler.ComandoProduction, p *production.Production) {
	p.ID = database.NewUUID()

	title := help.OR(cp.Metadata["Título Traduzido"], cp.Title)
	release := help.OR(cp.Metadata["Lançamento"], cp.Metadata["Ano de Lançamento"])
	p.Title = fmt.Sprintf("%s (%s)", title, release)

	p.Description = fmt.Sprintf(description_fmt, cp.Resume, metadataToMarkdownList(cp.Metadata))
	p.Genres = strings.ReplaceAll(cp.Metadata["Genêro"], " |", ", ")

	p.Images.Cover = cp.Cover
	p.Images.Trailer = strings.Split(cp.YoutubeTrailerEmbedURL, "/")[4]

	for _, magnet := range cp.Magnets {
		p.Downloads = append(p.Downloads, production.ProductionDownload{
			Name: magnet.Name, URL: magnet.URL,
		})
	}
}

func handleProductionImport(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get(r.FormValue("comando-url"))
	if err != nil {
		components.Warn(err.Error()).Render(r.Context(), w)
		return
	}
	defer res.Body.Close()

	var cp crawler.ComandoProduction
	err = crawler.ParseSingle(res.Body, &cp)
	if err != nil {
		components.Warn(err.Error()).Render(r.Context(), w)
		return
	}

	var p production.Production
	convertComandoProductionToProduction(&cp, &p)
	production.DoPostProcess(&p)
	p.Save()

	components.ProductionLanding(
		p, getCredential(r), components.ProductionRatingData{},
	).Render(r.Context(), w)
}
