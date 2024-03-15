package production

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/ikaio/tailmplx/database"
	"go.etcd.io/bbolt"
)

type Comment struct {
	ID            string
	AuthorName    string
	AuthorPicture string
	Content       string
}

func StoreComment(comment Comment, key ...string) error {
	return database.DB.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte("comments"))
		key := strings.Join(key, "-")

		b, err := json.Marshal(comment)
		if err != nil {
			return err
		}

		return bucket.Put([]byte(key), b)
	})
}

func GatherCommentsFor(prefix string) ([]*Comment, error) {
	list := make([]*Comment, 0, 50)
	database.DB.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("comments"))
		c := b.Cursor()

		for k, v := c.Seek([]byte(prefix)); k != nil && len(list) < cap(list) && bytes.HasPrefix(k, []byte(prefix)); k, v = c.Next() {
			var comment Comment
			err := json.Unmarshal(v, &comment)
			if err != nil {
				continue
			}
			list = append(list, &comment)
		}

		return nil
	})
	return list, nil
}
