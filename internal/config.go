package internal

import (
	"github.com/spf13/viper"
)

type Config struct {
	ViperInstance *viper.Viper
	AbsolutePath  string
	BookMarks     map[string]*BookMark `json:"bookmarks"`
	LastOpened    string               `json:"lastopened"`
}
