package internal

import (
	"encoding/json"
	"go.etcd.io/bbolt"
)

func (c *Config) WriteKey(bm BookMark) error {
	db, err := bbolt.Open(c.AbsolutePath, 0600, nil)
	if err != nil {
		panic(err)
	}
	err = db.Update(func(tx *bbolt.Tx) error {
		bookMarksBucket, err := tx.CreateBucketIfNotExists([]byte("BookMarks"))
		if err != nil {
			return err
		}
		buf, err := json.Marshal(bm)
		if err != nil {
			return err
		}
		bookMarksBucket.Put([]byte(bm.Name), buf)
		return nil
	})
	return nil
}

func (c *Config) GetBookMark(key string) (*BookMark, error) {
	db, err := bbolt.Open(c.AbsolutePath, 0600, &bbolt.Options{ReadOnly: true})
	if err != nil {
		panic(err)
	}
	bmToReturn := &BookMark{}
	err = db.View(func(tx *bbolt.Tx) error {
		bookMarksBucket := tx.Bucket([]byte("BookMarks"))
		if bookMarksBucket == nil {
			return err
		}
		buf := bookMarksBucket.Get([]byte(key))
		err := json.Unmarshal(buf, bmToReturn)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return bmToReturn, nil
}
