package models

import (
	"encoding/json"
	"time"

	"github.com/ikaio/tailmplx/database"
	"go.etcd.io/bbolt"
)

type Upload struct {
	ID     string
	Title  string
	Author string
	At     time.Time

	Files []string
}

func (c *Upload) Save() error {
	return database.DB.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("uploads"))

		buf, err := json.Marshal(c)
		if err != nil {
			return err
		}

		return b.Put([]byte(c.ID), buf)
	})
}

func ListUploads(limit int) (list []*Upload, err error) {
	err = database.DB.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("uploads"))
		c := b.Cursor()

		for k, v := c.First(); k != nil && len(list) < limit; k, v = c.Next() {
			var upload Upload
			json.Unmarshal(v, &upload)
			list = append(list, &upload)
		}

		return nil
	})
	return
}
