package crawler

import (
	"io"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type ComandoProduction struct {
	Title                  string
	URL                    string
	Cover                  string
	IMDBRate               string
	Resume                 string
	YoutubeTrailerEmbedURL string
	Magnets                []ComandoMagnet
	Metadata               map[string]string
}

type ComandoMagnet struct {
	Name string
	URL  string
}

func ParseSingle(r io.Reader, v *ComandoProduction) error {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return err
	}

	article := doc.Find("article").First()
	fromArticle(article, v)

	return nil
}

func fromArticle(article *goquery.Selection, v *ComandoProduction) {
	title := article.Find("header a")
	v.Title = title.Text()
	v.URL = title.AttrOr("href", "")

	content := article.Find(".entry-content")
	v.Cover = content.Find("img").AttrOr("src", "")
	v.IMDBRate = content.Find("a[href^=\"https://www.imdb.com\"]").Text()
	v.YoutubeTrailerEmbedURL = content.Find("iframe").AttrOr("src", "")

	paragraphs := content.Find("p")
	magnetanchors := content.Find("a[href^=\"magnet:\"]")

	paragraphs.Each(func(i int, p *goquery.Selection) {
		text := p.Text()
		if strings.Contains(text, "INFORMAÇÕES") {
			v.Metadata = make(map[string]string)
			parseMetadata(text, v.Metadata)
			return
		} else if strings.HasPrefix(text, "SINOPSE:") {
			v.Resume = text
			return
		}
	})

	magnetanchors.Each(func(i int, s *goquery.Selection) {
		url := s.AttrOr("href", "")
		name, err := nameFromMagnetHashURL(url)
		if err == nil {
			v.Magnets = append(v.Magnets, ComandoMagnet{Name: name, URL: url})
		}
	})
}

func parseMetadata(raw string, out map[string]string) {
	var clip string    // for storing
	var key string     // current key being set
	var expecting bool // if it is past a ':' and is expecting a value to a current key

loop:
	for _, r := range raw {
		switch r {
		case '\n':
			if expecting {
				out[key] = strings.TrimSpace(clip)
				expecting = false
			}
			clip = ""
		case ':':
			if expecting {
				clip += string(r)
				continue loop
			}
			key = strings.TrimSpace(clip)
			expecting = true
			clip = ""

		default: // store useful text
			clip += string(r)
		}
	}
}

func nameFromMagnetHashURL(raw string) (string, error) {
	URL, err := url.Parse(raw)
	if err != nil {
		return "", err
	}
	name := strings.ReplaceAll(URL.Query().Get("dn"), ".", " ")
	return name, nil
}
