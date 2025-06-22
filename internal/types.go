package internal

import "fmt"

type BookMark struct {
	Name string `mapstructure:"name"`
	Link string `mapstructure:"link"`
}

func (b BookMark) Println() {
	fmt.Println("### " + b.Name + " ###")
	fmt.Println("| " + b.Link)
}
