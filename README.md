# `gbox` - :inbox_tray: use github as storages :outbox_tray:

[![Coverage Status](https://coveralls.io/repos/github/zezaeoh/gbox/badge.svg?branch=main)](https://coveralls.io/github/zezaeoh/gbox?branch=main) [![Go Report Card](https://goreportcard.com/badge/github.com/zezaeoh/gbox)](https://goreportcard.com/report/github.com/zezaeoh/gbox)

## Installation

To download the latest release, run:

```bash
curl --silent --location "https://github.com/zezaeoh/gbox/releases/latest/download/gbox_$(uname -s)_amd64.tar.gz" | tar xz -C /tmp
sudo mv /tmp/gbox /usr/local/bin
```

For ARM system, please change ARCH (e.g. armv6, armv7 or arm64) accordingly

```bash
curl --silent --location "https://github.com/zezaeoh/gbox/releases/latest/download/gbox_$(uname -s)_arm64.tar.gz" | tar xz -C /tmp
sudo mv /tmp/gbox /usr/local/bin
```

macOS users can use [Homebrew](https://brew.sh):

```bash
brew tap zezaeoh/gbox
brew install gbox
```
