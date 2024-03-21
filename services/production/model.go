package production

import (
	"encoding/json"
	"errors"

	"github.com/ikaio/tailmplx/services/database"
	"go.etcd.io/bbolt"
)

var ProductionCache = make(map[string]*Production)

type ProductionImagesExtra struct {
	Url string `json:"url"`
}

type ProductionImages struct {
	Trailer string                  `json:"trailer"`
	Cover   string                  `json:"cover"`
	Banner  string                  `json:"banner"`
	Extras  []ProductionImagesExtra `json:"extras"`
}

type ProductionDownload struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type ProductionPostProcess struct {
	PrimaryColor string `json:"primary_color"`
	DarkerColor  string `json:"darker_color"`
	LigherColor  string `json:"ligher_color"`
}

type Production struct {
	ID          string                `json:"id"`
	Title       string                `json:"title"`
	Description string                `json:"description"`
	Genres      string                `json:"genres"`
	Images      ProductionImages      `json:"images"`
	Downloads   []ProductionDownload  `json:"downloads"`
	PostProcess ProductionPostProcess `json:"post_process"`
}

func (production *Production) Save() error {
	return database.DB.Update(func(transaction *bbolt.Tx) error {
		bucket := transaction.Bucket([]byte("productions"))

		buf, err := json.Marshal(production)
		if err != nil {
			return err
		}

		return bucket.Put([]byte(production.ID), buf)
	})
}

func FetchProduction(id string) (p *Production, err error) {

	err = database.DB.View(func(transaction *bbolt.Tx) error {
		bucket := transaction.Bucket([]byte("productions"))

		buf := bucket.Get([]byte(id))
		if buf == nil {
			return errors.New("could not get key: " + id)
		}

		var production Production
		err := json.Unmarshal(buf, &production)
		if err != nil {
			return err
		}

		p = &production
		return nil
	})

	return
}

func GetById(id string) (*Production, error) {
	if cached := ProductionCache[id]; cached != nil {
		return cached, nil
	}
	fetched, err := FetchProduction(id)
	if err != nil {
		return nil, err
	}
	ProductionCache[id] = fetched
	return fetched, nil
}

func ListProductions(limit int) (list []*Production, err error) {
	err = database.DB.View(func(transaction *bbolt.Tx) error {
		bucket := transaction.Bucket([]byte("productions"))
		cursor := bucket.Cursor()

		for k, v := cursor.First(); k != nil && len(list) < limit; k, v = cursor.Next() {
			var production Production
			json.Unmarshal(v, &production)
			list = append(list, &production)

			// save to cache if it ain't already
			if _, ok := ProductionCache[production.ID]; !ok {
				ProductionCache[production.ID] = &production
			}
		}

		return nil
	})
	return
}
