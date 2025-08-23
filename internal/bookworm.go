package internal

import (
	"errors"
)

type BookWorm struct {
	Cfg       *Config
	BookMarks map[string]*BookMark
}

func Init() *BookWorm {
	cfg, err := getConfig()
	if err != nil {
		panic(err)
	}
	return &BookWorm{
		Cfg:       cfg,
		BookMarks: cfg.BookMarks,
	}
}

// Registister Config writes all of the changes to the Config
// This is a little janky, but oh well
func (w *BookWorm) RegisterConfig() error {
	return w.Cfg.writeConfig()
}

func (w *BookWorm) SetLastOpened(bm BookMark) error {
	w.Cfg.LastOpened = bm.Link
	return w.RegisterConfig()
}

func (w *BookWorm) SetTags(name string, tags []string) error {
	bm, ok := w.Cfg.BookMarks[name]
	if !ok {
		return errors.New("Bookmark not in mealhouse")
	}
	bm.Tags = append(bm.Tags, tags...)
	// Rewriting all of the bookmarks each time is not great
	return w.RegisterConfig()
}

func (w *BookWorm) NewBookMark(name string, link string, tags []string) error {
	w.Cfg.BookMarks[name] = &BookMark{
		Name: name,
		Link: link,
		Tags: tags,
	}
	return w.RegisterConfig()
}

func (w *BookWorm) DeleteBookMark(name string) error {
	w.Cfg.BookMarks[name] = &BookMark{}
	delete(w.Cfg.BookMarks, name)
	return w.RegisterConfig()
}
