package state

import (
	"strings"

	"go.etcd.io/bbolt"
)

type State struct {
	AbsolutePath string
	Links        map[string]string
	Tags         map[string][]string
}

func (c *State) Dump(key string) error {
	db, err := bbolt.Open(c.AbsolutePath, 0600, nil)
	if err != nil {
		panic(err)
	}
	err = db.Update(func(tx *bbolt.Tx) error {
		LinksBucket, err := tx.CreateBucketIfNotExists([]byte("Links"))
		if err != nil {
			panic(err)
		}
		TagsBucket, err := tx.CreateBucketIfNotExists([]byte("Tags"))
		if err != nil {
			panic(err)
		}
		for k, v := range c.Links {
			LinksBucket.Put([]byte(k), []byte(v))
			tags, ok := c.Tags[k]
			if !ok {
				continue
			}
			tagMash := ""
			for _, tag := range tags {
				if !strings.Contains(tag, ";") {
					tagMash += tag + ";"
				}
			}
			TagsBucket.Put([]byte(k), []byte(tagMash))
		}
		return nil
	})
	return nil
}
