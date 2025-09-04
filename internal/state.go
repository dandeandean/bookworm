package internal

import (
	"encoding/json"
	"errors"
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
		panic(err)
	}
	defer db.Close()
	buf, err := json.Marshal(bm)
	if err != nil {
		panic(err)
	}
	err = db.Update(func(tx *bbolt.Tx) error {
		bookMarksBucket, err := tx.CreateBucketIfNotExists([]byte("bookmarks"))
		if err != nil {
			panic(err)
		}
		return bookMarksBucket.Put([]byte(bm.Name), buf)
	})
	return nil
}

func (bw *BookWorm) enumBookMarks(key string) ([]*BookMark, error) {
	return nil, nil
}

func (bw *BookWorm) getBookMark(key string) (*BookMark, error) {
	db, err := bbolt.Open(bw.Cfg.DbPath, 0600, &bbolt.Options{ReadOnly: true, Timeout: time.Second})
	if err != nil {
		panic(err)
	}
	defer db.Close()
	bmToReturn := &BookMark{}
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
	err = json.Unmarshal(buf, bmToReturn)
	if err != nil {
		return nil, err
	}
	return bmToReturn, nil
}
