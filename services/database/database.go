package database

import (
	"encoding/binary"

	"github.com/bwmarrin/snowflake"
	"go.etcd.io/bbolt"
)

var DB *bbolt.DB
var SF *snowflake.Node

func Setup(database_path string) {

	sf, err := snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}

	db, err := bbolt.Open(database_path, 0600, nil)
	if err != nil {
		panic(err)
	}

	err = db.Update(func(tx *bbolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte("productions"))
		tx.CreateBucketIfNotExists([]byte("ratings"))
		tx.CreateBucketIfNotExists([]byte("comments"))
		return nil
	})

	if err != nil {
		panic(err)
	}

	DB = db
	SF = sf
}

func NewUUID() string {
	return SF.Generate().String()
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
