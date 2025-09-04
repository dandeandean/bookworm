package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"go.etcd.io/bbolt"
)

func (bw *BookWorm) writeBookMark(key string) error {
	bm := bw.BookMarks[key]
	if bm == nil {
		return errors.New("BookMark is Nil")
	}
	db, err := bbolt.Open(bw.Cfg.DbPath, 0600, &bbolt.Options{Timeout: time.Second})
	if err != nil {
		return err
	}
	defer db.Close()
	buf, err := json.Marshal(bm)
	if err != nil {
		return err
	}
	err = db.Update(func(tx *bbolt.Tx) error {
		bookMarksBucket, err := tx.CreateBucketIfNotExists([]byte("bookmarks"))
		if err != nil {
			return err
		}
		return bookMarksBucket.Put([]byte(bm.Name), buf)
	})
	return nil
}

func (c *Config) dumpBookMarks() (map[string]*BookMark, error) {
	db, err := bbolt.Open(c.DbPath, 0600, &bbolt.Options{Timeout: time.Second})
	if err != nil {
		return nil, err
	}
	defer db.Close()
	bmsRaw := make(map[string][]byte)
	db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("bookmarks"))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			bmsRaw[string(k)] = v
		}
		return nil
	})
	bookmarks := make(map[string]*BookMark)
	for k, v := range bmsRaw {
		b, err := bytesToBookMark(v)
		if err != nil {
			return nil, err
		}
		bookmarks[k] = b
	}
	return bookmarks, nil
}

func (bw *BookWorm) getBookMark(key string) (*BookMark, error) {
	db, err := bbolt.Open(bw.Cfg.DbPath, 0600, &bbolt.Options{ReadOnly: true, Timeout: time.Second})
	if err != nil {
		panic(err)
	}
	defer db.Close()
	var buf []byte
	err = db.View(
		func(tx *bbolt.Tx) error {
			bookMarksBucket := tx.Bucket([]byte("bookmarks"))
			if bookMarksBucket == nil {
				return err
			}
			buf = bookMarksBucket.Get([]byte(key))
			return nil
		})
	if err != nil {
		return nil, err
	}

	return bytesToBookMark(buf)
}
