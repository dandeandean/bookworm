package internal

import "fmt"

type BookMark struct {
	Name string   `json:"name"`
	Link string   `json:"link"`
	Tags []string `json:"tags"`
}

func (b BookMark) Println() {
	if len(b.Tags) != 0 {
		fmt.Println(b.Name, b.Tags)
	} else {
		fmt.Println(b.Name)
	}
}

func (b BookMark) Open() {
	OpenURL(b.Link)
}
