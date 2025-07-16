package internal

import (
	"errors"
	"fmt"
)

type BookWorm struct {
	Cfg *Config
}

type BookMark struct {
	Name string   `mapstructure:"name"`
	Link string   `mapstructure:"link"`
	Tags []string `mapstructure:"tags"`
}

func Init() *BookWorm {
	cfg := GetConfig()
	return &BookWorm{
		Cfg: cfg,
	}
}

func (w *BookWorm) RegisterConfig() error {
	// This is a little janky, but oh well
	return w.Cfg.ViperInstance.WriteConfig()
}

func (b BookMark) Println() {
	fmt.Println(b.Name + ": " + b.Link)
	if len(b.Tags) != 0 {
		fmt.Print(" ")
		fmt.Println(b.Tags)
	}
}

func (w *BookWorm) SetLastOpened(bm BookMark) error {
	w.Cfg.LastOpened = bm.Link
	w.Cfg.ViperInstance.Set("lastopened", bm.Link)
	return w.RegisterConfig()
}

func (w *BookWorm) SetTags(name string, tags []string) error {
	bm, ok := w.Cfg.BookMarks[name]
	if !ok {
		return errors.New("Bookmark not in mealhouse")
	}
	bm.Tags = append(bm.Tags, tags...)
	// Rewriting all of the bookmarks each time is not great
	w.Cfg.ViperInstance.Set("bookmarks", w.Cfg.BookMarks)
	return w.RegisterConfig()
}

func (w *BookWorm) NewBookMark(name string, link string, tags []string) error {
	w.Cfg.BookMarks[name] = &BookMark{
		Name: name,
		Link: link,
		Tags: tags,
	}
	w.Cfg.ViperInstance.Set("bookmarks", w.Cfg.BookMarks)
	return w.RegisterConfig()
}

func (w *BookWorm) DeleteBookMark(name string) error {
	w.Cfg.BookMarks[name] = &BookMark{}
	delete(w.Cfg.BookMarks, name)
	w.Cfg.ViperInstance.Set("bookmarks", w.Cfg.BookMarks)
	return w.RegisterConfig()
}
