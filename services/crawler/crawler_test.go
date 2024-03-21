package crawler

import (
	"io"
	"net/url"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFullComandoProductionParsing(t *testing.T) {

	tests := []struct {
		file     string
		title    string
		resume   string
		metadata map[string]string
		magnets  []ComandoMagnet
	}{
		/* { // new production format
			file:     "./testdata/argylle.html",
			title:    "Argylle: O Superespião",
			resume:   "Argylle – O Superespião é um filme de comédia e ação",
			metadata: map[string]string{"Título Traduzido": "Argylle: O Superespião"},
			magnets:  []ComandoMagnet{{Name: "WEB-DL 1080p Dual Áudio 5.1 (MKV) | (3.30 GB)", URL: "magnet:?xt=urn:btih:a4eb66d"}},
		},
		{ // slight older format
			file:     "./testdata/angie.html",
			title:    "Angie: Garotas Perdidas Torrent",
			resume:   "SINOPSE: Angie é uma adolescente",
			metadata: map[string]string{"Título Traduzido": "Angie: Garotas Perdidas"},
			magnets:  []ComandoMagnet{{Name: "BluRay 1080p Dual Áudio (MKV) (1.87 GB)", URL: "magnet:?xt=urn:btih:3722cb6bdf"}},
		}, */
		{
			file:     "./testdata/regime.html",
			title:    "O Regime",
			resume:   "Na minissérie The Regime, a",
			metadata: map[string]string{"Ano de Lançamento": "2024"},
			magnets:  []ComandoMagnet{{Name: "O Regime S01E01 Dia da Vitoria 1080p WEB-DL DUAL", URL: "magnet:?xt=urn:btih:8d2e0f7f"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			file, err := os.Open(tt.file)
			assert.Nil(t, err)

			var p ComandoProduction
			err = ParseSingle(file, &p)
			assert.Nil(t, err)

			assert.Contains(t, p.Title, tt.title)
			assert.Contains(t, p.Resume, tt.resume)

			_, err = url.Parse(p.Cover)
			assert.Nil(t, err)

			for k, ev := range tt.metadata {
				assert.Equal(t, ev, p.Metadata[k])
			}

			for i, em := range tt.magnets {
				assert.Equal(t, em.Name, p.Magnets[i].Name)
				assert.Contains(t, p.Magnets[i].URL, em.URL)
			}
		})
	}
}

func TestComandoProductionMetadataParsing(t *testing.T) {
	contentinfotext, err := os.Open("./testdata/contentinfofield.txt")
	assert.Nil(t, err)

	b, err := io.ReadAll(contentinfotext)
	assert.Nil(t, err)

	out := make(map[string]string)
	parseMetadata(string(b), out)
	assert.Nil(t, err)

	assert.Equal(t, "MKV", out["Formato"])
	assert.Equal(t, "Torrent", out["Servidor"])
}

func TestMagnetNameFromURL(t *testing.T) {
	name, err := nameFromMagnetHashURL("magnet:?xt=urn:btih:8d2e0f7fc71a69911ebd05361e3a1b688feb8b84&dn=O.Regime.S01E01.Dia.da.Vitoria.1080p.WEB-DL.DUAL")
	assert.Nil(t, err)
	assert.Equal(t, "O Regime S01E01 Dia da Vitoria 1080p WEB-DL DUAL", name)
}
