package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/go-yaml/yaml"
	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/pkg/browser"
)

const (
	configDir        = ".config/fzf-bookmark-opener"
	configFileName   = "config.yaml"
	defaultBookmarks = `bookmarks:
  - title: 'fzf-bookmark-opener(You can edit bookmarks to edit "~/.config/fzf-bookmark-opener/config.yaml"!)'
    url: 'https://github.com/kyu08/fzf-bookmark-opener'
  - title: 'fzf-bookmark-opener Issues'
    url: 'https://github.com/kyu08/fzf-bookmark-opener/issues'
  - title: 'fzf-bookmark-opener PRs'
    url: 'https://github.com/kyu08/fzf-bookmark-opener/pulls'
`
)

type Config struct {
	Config []Bookmark `yaml:"bookmarks"`
}

type Bookmark struct {
	Title string `yaml:"title"`
	Url   string `yaml:"url"`
}

func main() {
	// initで設定ファイルの作成
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprint(os.Stderr, "err")
	}
	initialize(homeDir)

	// 設定ファイルの読み込み
	bookmarks, err := loadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// fzfで選択する
	index, err := find(bookmarks)
	if err != nil {
		return
	}

	// ブラウザで表示する
	browser.OpenURL(bookmarks[index].Url)
}

// initialize テスト容易性のために設定ファイルを配置するディレクトリを引数で受け取っている
func initialize(workingDir string) error {
	configDirFullPath := path.Join(workingDir, configDir)

	// 設定ファイル用ディレクトリの作成
	if _, err := os.Stat(configDirFullPath); err != nil {
		if err := os.MkdirAll(configDirFullPath, 0o755); err != nil {
			return fmt.Errorf("fail: os.MkdirAll: %w", err)
		}
		fmt.Printf("~%s created.\n", configDirFullPath)
	}

	// 設定ファイルの作成
	configFileFullPath := path.Join(configDirFullPath, configFileName)
	if _, err := os.Stat(configFileFullPath); err != nil {
		// ファイルが存在していない場合は作成する
		if err := os.WriteFile(configFileFullPath, []byte(defaultBookmarks), 0o777); err != nil {
			return fmt.Errorf("os.WriteFile: %w", err)
		}
		fmt.Printf("~%s created.\n", configFileFullPath)
		fmt.Println("You can add, update, delete bookmarks to edit `config.yaml`.")
	}

	return nil
}

func loadConfig() ([]Bookmark, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("loadConfig os.UserHomeDir err:", err)
	}

	configFilePath := filepath.Join(homeDir, configDir, configFileName)
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
