<div align="center">

# fzf-bookmark-opener

`fzf-bookmark-opener` is the command line tool that opens bookmark using fzf.

![Language:Go](https://img.shields.io/static/v1?label=Language&message=Go&color=blue&style=flat-square)
![License:MIT](https://img.shields.io/static/v1?label=License&message=MIT&color=blue&style=flat-square)
[![Latest Release](https://img.shields.io/github/v/release/kyu08/fzf-bookmark-opener?style=flat-square)](https://github.com/kyu08/fzf-bookmark-opener/releases/latest)

</div>

![how to use](https://user-images.githubusercontent.com/49891479/218272272-e693c10d-c810-458a-bf46-9c3a4a2fe45a.gif)

# Installation
🚨 This command tested only on macOS.

```
go install github.com/kyu08/fzf-bookmark-opener@latest
```

# How it works
- This command generates the configuration file(`~/.config/fzf-bookmark-opener/config.yaml`) when initial launch.
- `config.yaml` is like below.
- You can add, update, delete bookmarks to edit this file.

```yaml
bookmarks:
  - title: 'fzf-bookmark-opener'
    url: 'https://github.com/kyu08/fzf-bookmark-opener'
  - title: 'fzf-bookmark-opener Issues'
    url: 'https://github.com/kyu08/fzf-bookmark-opener/issues'
  - title: 'fzf-bookmark-opener PRs'
    url: 'https://github.com/kyu08/fzf-bookmark-opener/pulls'
```

