package main

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_initialize(t *testing.T) {
	t.Run("initialize", func(t *testing.T) {
		tmpDir, err := os.MkdirTemp("", "test")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(tmpDir)

		if err = initialize(tmpDir); err != nil {
			t.Fatal(err)
		}

		configDir := path.Join(tmpDir, ".config", "fzf-bookmark-opener")
		assert.DirExists(t, configDir)

		filePath := path.Join(configDir, "config.yaml")
		assert.FileExists(t, filePath)

		b, err := ioutil.ReadFile(filePath)
		assert.NoError(t, err)

		content := string(b)
		assert.Contains(t, content, "bookmarks")
		assert.Contains(t, content, "title")
		assert.Contains(t, content, "url")
	})
}
