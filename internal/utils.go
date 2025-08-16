package internal

import (
	"errors"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/spf13/viper"
)

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

// openURL opens the specified URL in the default browser of the user.
// From https://gist.github.com/sevkin/9798d67b2cb9d07cb05f89f14ba682f8
func OpenURL(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "rundll32"
		args = []string{"url.dll,FileProtocolHandler", url}
	case "darwin":
		cmd = "open"
		args = []string{url}
	default:
		if isWSL() {
			cmd = "cmd.exe"
			args = []string{"/c", "start", url}
		} else {
			cmd = "xdg-open"
			args = []string{url}
		}
	}
	if len(args) > 1 {
		// args[0] is used for 'start' command argument, to prevent issues with URLs starting with a quote
		args = append(args[:1], append([]string{""}, args[1:]...)...)
	}
	return exec.Command(cmd, args...).Start()
}

// isWSL checks if the Go program is running inside Windows Subsystem for Linux
// From https://gist.github.com/sevkin/9798d67b2cb9d07cb05f89f14ba682f8
func isWSL() bool {
	releaseData, err := exec.Command("uname", "-r").Output()
	if err != nil {
		return false
	}
	return strings.Contains(strings.ToLower(string(releaseData)), "microsoft")
}

// This is going to need a lot of work
func IsValidUrl(url string) bool {
	return strings.Contains(url, "https://") ||
		strings.Contains(url, "http://")
}
