package main

import (
	"log"

	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/pkg/browser"
)

type Bookmark struct {
	Title string
	Url   string
}

func main() {
	// TODO: 設定ファイルを読み込む
	var bookmarks = []Bookmark{
		{"fzf-bookmark-opener hoge", "https://github.com/kyu08/fzf-bookmark-opener"},
		{"blog", "https://github.com/kyu08/blog"},
	}

	// fzfで選択する
	index, err := find()
	if err != nil {
		log.Fatal(err)
	}

	// ブラウザで表示する
	browser.OpenURL(bookmarks[index].Url)
}

func find() (int, error) {
	return fuzzyfinder.Find(
		bookmarks,
		func(i int) string {
			return bookmarks[i].Title
		},
	)
}
