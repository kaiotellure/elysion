package database

import (
	"encoding/binary"
	"encoding/json"

	"github.com/bwmarrin/snowflake"
	"go.etcd.io/bbolt"
)

var DB *bbolt.DB
var SnowflakeNode *snowflake.Node

type Upload struct {
	ID       string
	Filename string
	Title    string
	Author   string
}

func (c *Upload) Save() error {
	return DB.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("uploads"))

		buf, err := json.Marshal(c)
		if err != nil {
			return err
		}

		return b.Put([]byte(c.ID), buf)
	})
}

func ListUploads(limit int) (list []*Upload, err error) {
	err = DB.View(func(tx *bbolt.Tx) error {
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

func Setup(database_path string) {

	// Setup snowflake node for id generation
	node, err := snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}
	SnowflakeNode = node

	db, err := bbolt.Open(database_path, 0600, nil)
	if err != nil {
		panic(err)
	}

	err = db.Update(func(tx *bbolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte("uploads"))
		return nil
	})

	if err != nil {
		panic(err)
	}

	DB = db
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
