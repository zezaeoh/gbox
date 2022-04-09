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

## Basic Usage

### Add storage config

```bash
gbox storage add
```

```bash
$ gbox storage add
? Name of storage gbox-test
? Kind of storage github
? Repository URL https://github.com/zezaeoh/gbox-test
? Repository Branch main
? Authentication method none
* Successfully add github storage config: gbox-test
```

after simple survey, storage config will be store in your local path `~/.config/gbox/`

if you want to add private repository, you should add github token to your storage config

### List storage configs

```bash
gbox storage list
```

```bash
$ gbox storage list
* gbox-test
* ✓ gbox-storage
```

active storage is checked

### Set active storage

```bash
gbox storage set <name>
```

```bash
$ gbox storage set gbox-test
* Storage Configured: gbox-test

# check active storage
$ gbox storage list
* ✓ gbox-test
* gbox-storage
```

### Set data

```bash
gbox set <name> <data>
```

```bash
$ gbox set test/my-secret supersupersecret
* Set: test/my-secret
```

### Get data

```bash
gbox get <name>
```

```bash
$ gbox get test/my-secret
supersupersecret
```

### Delete data

```bash
gbox delete <name>
```

```bash
$ gbox delete test/my-secret
* Delete: test/my-secret
```

### List data
```bash
gbox list
```

```bash
$ gbox list
/
├─ test/
│  ├─ sample/
│  │  └─ whoami
│  ├─ test
│  └─ test2
└─ github/
   └─ token/
      ├─ zezaeoh
      └─ zezaeoh2
```
