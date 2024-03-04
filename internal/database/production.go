package database

import (
	"encoding/json"
	"errors"

	"go.etcd.io/bbolt"
)

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

type ProductionProperties struct {
	PrimaryColor                 string `json:"primary_color"`
	DarkerColor                  string `json:"darker_color"`
	LigherColor                  string `json:"ligher_color"`
	PostProcessedDescriptionHTML string `json:"post_processed_description_html"`
}

type Production struct {
	ID          string               `json:"id"`
	Title       string               `json:"title"`
	Description string               `json:"description"`
	Genres      string               `json:"genres"`
	Images      ProductionImages     `json:"images"`
	Downloads   []ProductionDownload `json:"downloads"`
	Properties  ProductionProperties `json:"properties"`
}

func (production *Production) Save() error {
	return DB.Update(func(transaction *bbolt.Tx) error {
		bucket := transaction.Bucket([]byte("productions"))

		buf, err := json.Marshal(production)
		if err != nil {
			return err
		}

		return bucket.Put([]byte(production.ID), buf)
	})
}

func GetProduction(id string) (prod Production, err error) {
	DB.View(func(transaction *bbolt.Tx) error {
		bucket := transaction.Bucket([]byte("productions"))

		buf := bucket.Get([]byte(id))
		if buf == nil {
			err = errors.New("could not get key: " + id)
			return err
		}

		err = json.Unmarshal(buf, &prod)
		if err != nil {
			return err
		}

		return err
	})
	return
}

func ListProductions(limit int) (list []*Production, err error) {
	err = DB.View(func(transaction *bbolt.Tx) error {
		bucket := transaction.Bucket([]byte("productions"))
		cursor := bucket.Cursor()

		for k, v := cursor.First(); k != nil && len(list) < limit; k, v = cursor.Next() {
			var production Production
			json.Unmarshal(v, &production)
			list = append(list, &production)
		}

		return nil
	})
	return
}
