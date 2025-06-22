package config

import (
	"errors"
	"os"

	. "github.com/dandeandean/bookworm/internal"
	"github.com/spf13/viper"
)

type Config struct {
	ViperInstance *viper.Viper
	BookMarks     []BookMark `mapstructure:"bookmarks"`
	LastOpened    string     `mapstructure:"lastopened"`
}

func (c *Config) NewBookMark(name string, link string) {
	c.BookMarks = append(c.BookMarks,
		BookMark{
			Name: name,
			Link: link,
		},
	)
	c.ViperInstance.Set("bookmarks", c.BookMarks)
	err := c.ViperInstance.WriteConfig()
	if err != nil {
		panic(err)
	}
}

func GetConfig() *Config {
	homedir := os.Getenv("HOME")
	var cfg Config
	cfg.ViperInstance = viper.New()
	cfg.ViperInstance.SetConfigType("yaml")
	cfg.ViperInstance.SetConfigName("bookworm")
	cfg.ViperInstance.AddConfigPath(homedir + "/.config/")
	cfg.ViperInstance.SetDefault("bookmarks", []BookMark{})
	if !IsConfigThere() {
		err := WriteNewConfig()
		if err != nil {
			panic("Error Writing Config")
		}
	}
	cfg.ViperInstance.ReadInConfig()
	err := cfg.ViperInstance.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	return &cfg
}

func IsConfigThere() bool {
	homedir := os.Getenv("HOME")
	_, try := os.Stat(homedir + "/.config/bookworm.yml")
	_, try1 := os.Stat(homedir + "/.bookworm.yml")
	if os.IsNotExist(try) && os.IsNotExist(try1) {
		return false
	}
	return true
}

func WriteNewConfig() error {
	homedir := os.Getenv("HOME")
	if IsConfigThere() {
		return errors.New("config is already there!")
	}
	configInfo, err := os.Stat(homedir + "/.config")
	if err != nil {
		panic(err)
	}
	if !configInfo.IsDir() {
		return errors.New("~/.config is not a directory!")
	}
	viper.SetDefault("bookmarks", []BookMark{})
	viper.SetDefault("lastopened", "")
	viper.WriteConfigAs(homedir + "/.config/bookworm.yml")
	return nil
}
