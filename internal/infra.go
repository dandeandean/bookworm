package internal

import (
	"fmt"
)

type BookWorm struct {
	Cfg *Config
}

type BookMark struct {
	Name string `mapstructure:"name"`
	Link string `mapstructure:"link"`
}

func Init() *BookWorm {
	cfg := GetConfig()
	return &BookWorm{
		Cfg: cfg,
	}
}

func (b BookMark) Println() {
	fmt.Println("######### " + b.Name + " #########")
	fmt.Println("| " + b.Link)
}

func (w *BookWorm) NewBookMark(name string, link string) {
	w.Cfg.BookMarks[name] = BookMark{
		Name: name,
		Link: link,
	}
	w.Cfg.ViperInstance.Set("bookmarks", w.Cfg.BookMarks)
	err := w.Cfg.ViperInstance.WriteConfig()
	if err != nil {
		panic(err)
	}
}
