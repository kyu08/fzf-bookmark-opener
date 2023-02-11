package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/go-yaml/yaml"
	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/pkg/browser"
)

func main() {
	// TODO: initで設定ファイルの作成
	// 設定ファイルの読み込み
	bookmarks, err := loadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// fzfで選択する
	// TODO: escで終了した時に何も表示しないようにする
	index, err := find(bookmarks)
	if err != nil {
		log.Fatal(err)
	}

	// ブラウザで表示する
	browser.OpenURL(bookmarks[index].Url)
}

type Config struct {
	Config []Bookmark `yaml:"bookmarks"`
}

type Bookmark struct {
	Title string `yaml:"title"`
	Url   string `yaml:"url"`
}

func loadConfig() ([]Bookmark, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("loadConfig os.UserHomeDir err:", err)
	}

	configFilePath := filepath.Join(homeDir, ".config", "fzf-bookmark-opener", "config.yaml")
	b, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		log.Fatal("loadConfig ioutil.ReadFile err:", err)
	}

	var c Config
	if err := yaml.Unmarshal(b, &c); err != nil {
		log.Fatal(err)
	}

	return c.Config, nil
}

func find(bookmarks []Bookmark) (int, error) {
	return fuzzyfinder.Find(
		bookmarks,
		func(i int) string {
			return bookmarks[i].Title
		},
	)
}
