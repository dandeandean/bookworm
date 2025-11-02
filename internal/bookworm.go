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
	cfg, err := getConfig("")
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
	cfg, err := initConfig("")
	if err != nil {
		return nil, err
	}
	if cfg == nil {
		return nil, err
	}
	return &BookWorm{
		Cfg:       cfg,
		BookMarks: make(map[string]*BookMark),
	}, nil
}

func (w *BookWorm) GetAllRaw() (map[string][]byte, error) {
	return readAll(w.Cfg.DbPath)
}

func (w *BookWorm) GetOneRaw(key string) ([]byte, error) {
	return readOne(w.Cfg.DbPath, key)
}

func (w *BookWorm) SetLastOpened(bm BookMark) error {
	w.Cfg.LastOpened = bm.Link
	return w.Cfg.writeConfig("")
}

func (w *BookWorm) SetTags(name string, tags []string) error {
	bm, ok := w.BookMarks[name]
	if !ok {
		return errors.New("Bookmark not in mealhouse")
	}
	bm.Tags = append(bm.Tags, tags...)
	// Rewriting all of the bookmarks each time is not great
	return w.writeBookMark(name)
}

func (w *BookWorm) NewBookMark(name string, link string, tags []string) error {
	w.BookMarks[name] = &BookMark{
		Name: name,
		Link: link,
		Tags: tags,
	}
	return w.writeBookMark(name)
}

func (w *BookWorm) DeleteBookMark(name string) error {
	w.BookMarks[name] = &BookMark{}
	delete(w.BookMarks, name)
	return w.deleteBookMark(name)
}

// TODO: change signature of this func to return error
func (w *BookWorm) GetBookMark(name string) *BookMark {
	bmRaw, err := w.GetOneRaw(name)
	if err != nil {
		printIfVerbose(err)
		return nil
	}
	bm, err := bytesToBookMark(bmRaw)
	if err != nil {
		printIfVerbose(err)
		return nil
	}
	return bm
}

// TODO: change signature of this func to return error
func (w *BookWorm) ListBookMarks(tagFilter string) []*BookMark {
	allBookMarks, err := w.Cfg.enumBookMarks()
	if err != nil {
		printIfVerbose(err)
		return nil
	}
	out := make([]*BookMark, 0)
	for _, b := range allBookMarks {
		if slices.Contains(b.Tags, tagFilter) || tagFilter == "" {
			out = append(out, b)
		}
	}
	return out
}

func (w *BookWorm) LenBookMarks() int {
	allBookMarks, err := w.Cfg.enumBookMarks()
	if err != nil {
		printIfVerbose(err)
		return 0
	}
	return len(allBookMarks)
}
