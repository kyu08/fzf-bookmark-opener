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

var bookmarks = []Bookmark{
	{"fzf-bookmark-opener hoge", "https://github.com/kyu08/fzf-bookmark-opener"},
	{"blog", "https://github.com/kyu08/blog"},
}

func main() {
	// TODO: 設定ファイルを読み込む
	// fzfで選択してもらう
	index, err := fuzzyfinder.Find(
		bookmarks,
		func(i int) string {
			return bookmarks[i].Title
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	// ブラウザで表示する
	browser.OpenURL(bookmarks[index].Url)
}
