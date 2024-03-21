package production

import (
	"bytes"
	"strings"

	"github.com/ikaio/tailmplx/services/database"
	"github.com/ikaio/tailmplx/help"
	"go.etcd.io/bbolt"
)

// key gets dash-joined: foo, bar = "foo-bar"
func StoreRating(value string, key ...string) error {
	return database.DB.Update(func(transaction *bbolt.Tx) error {
		bucket := transaction.Bucket([]byte("ratings"))
		return bucket.Put([]byte(strings.Join(key, "-")), []byte(value))
	})
}

// key gets dash-joined: foo, bar = "foo-bar"
func RetrieveRating(fallback string, key ...string) (value string) {
	database.DB.View(func(transaction *bbolt.Tx) error {
		bucket := transaction.Bucket([]byte("ratings"))

		b := bucket.Get([]byte(strings.Join(key, "-")))
		value = help.NZ(string(b), fallback)

		return nil
	})
	return
}

func CountProductionRating(id string) (int, int) {
	typemap := make(map[string]int)

	database.DB.View(func(transaction *bbolt.Tx) error {
		cursor := transaction.Bucket([]byte("ratings")).Cursor()
		prefix := []byte(id)

		for k, v := cursor.Seek(prefix); k != nil && bytes.HasPrefix(k, prefix); k, v = cursor.Next() {
			typemap[string(v)]++
		}

		return nil
	})

	return typemap["love"], typemap["like"]
}
