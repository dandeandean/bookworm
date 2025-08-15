package state

import "go.etcd.io/bbolt"

type State struct {
	AbsolutePath string
	Links        map[string]string
	Tags         map[string][]string
}

func (c *State) WriteLink(key string) error {
	db, err := bbolt.Open("my.db", 0600, nil)
	if err != nil {
		panic(err)
	}
	tx, err := db.Begin(true)
	return nil
}
