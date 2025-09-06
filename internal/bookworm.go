package internal

import (
	"errors"
	"slices"
)

type BookWorm struct {
	Cfg       *Config
	BookMarks map[string]*BookMark
}

// get already init'd config
func Get() (*BookWorm, error) {
	cfg, err := getConfig()
	if err != nil {
		return nil, err
	}
	bms, err := cfg.enumBookMarks()
	if err != nil {
		return nil, err
	}
	if bms == nil {
		return nil, errors.New("Bookmarks are nil!")
	}
	return &BookWorm{
		Cfg:       cfg,
		BookMarks: bms,
	}, nil
}

// Init a Config that is not already there
func Init() (*BookWorm, error) {
	cfg, err := initConfig()
	if err != nil {
		return nil, err
	}
	if cfg == nil {
		return nil, err
	}
	bms, err := cfg.enumBookMarks()
	if err != nil {
		return nil, err
	}
	if bms == nil {
		bms = make(map[string]*BookMark)
	}
	return &BookWorm{
		Cfg:       cfg,
		BookMarks: bms,
	}, nil
}

// Registister Config writes all of the changes to the Config
// This is a little janky, but oh well
func (w *BookWorm) RegisterConfig() error {
	// for now we'll store the bookmarks in the config
	// this should be replaced by a proper db
	w.Cfg.BookMarks = w.BookMarks
	return w.Cfg.writeConfig()
}

func (w *BookWorm) SetLastOpened(bm BookMark) error {
	w.Cfg.LastOpened = bm.Link
	return w.RegisterConfig()
}

func (w *BookWorm) SetTags(name string, tags []string) error {
	bm, ok := w.BookMarks[name]
	if !ok {
		return errors.New("Bookmark not in mealhouse")
	}
	bm.Tags = append(bm.Tags, tags...)
	// Rewriting all of the bookmarks each time is not great
	return w.RegisterConfig()
}

func (w *BookWorm) NewBookMark(name string, link string, tags []string) error {
	w.BookMarks[name] = &BookMark{
		Name: name,
		Link: link,
		Tags: tags,
	}
	w.writeBookMark(name)
	return w.RegisterConfig()
}

func (w *BookWorm) DeleteBookMark(name string) error {
	w.BookMarks[name] = &BookMark{}
	delete(w.BookMarks, name)
	return w.RegisterConfig()
}

func (w *BookWorm) GetBookMark(name string) *BookMark {
	return w.BookMarks[name]
}

func (w *BookWorm) ListBookMarks(tagFilter string) []*BookMark {
	out := make([]*BookMark, 0)
	for _, b := range w.BookMarks {
		if slices.Contains(b.Tags, tagFilter) || tagFilter == "" {
			out = append(out, b)
		}
	}
	return out
}

func (w *BookWorm) LenBookMarks() int {
	return len(w.BookMarks)
}
