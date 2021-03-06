# fetch-github-ssh-key

* GitHubの任意ユーザのSSH公開鍵を取得し、ペーストで .ssh/authorized_keys に設定するツール
* -g オプションをつけると、authorized_keysに追記するスクリプトを生成します。

# インストールとビルド

## インストール

最新のバージョン番号(TARGET_VERSION)は[こちらで](https://github.com/mamemomonga/github-sshkey-fetch/releases)確認してください。

macOSの例

	$ export TARGET_VERSION=v0.0.3
	$ curl -Lo ./github-sshkey-fetch https://github.com/mamemomonga/github-sshkey-fetch/releases/download/$TARGET_VERSION/github-sshkey-fetch-darwin-amd64
	$ chmod 755 ./github-sshkey-fetch
	$ mkdir ~/bin
	$ mv ./github-sshkey-fetch ~/bin/
	$ echo 'export PATH="$HOME/bin:$PATH"' >> ~/.bashrc
	$ exec $SHELL
	$ github-sshkey-fetch -v

## ソースでの実行

	$ go run ./src

## ビルド

	$ make

## Dockerコンテナを使い各アーキテクチャ向けにビルド

	$ make docker

# 使い方

	$ fetch-github-ssh-key

### SSH公開鍵を表示する

	$ fetch-github-ssh-key -u mamemomonga

カンマ区切りで複数ユーザも対応可能

	$ fetch-github-ssh-key -u mamemomonga,user1,user2

### SSH公開鍵を設定するコマンドを表示する

	$ fetch-github-ssh-key -u mamemomonga -g

	[ debug ] 2020/06/04 05:54:43 main.go:30: Username: mamemomonga
	[ debug ] 2020/06/04 05:54:43 fetch.go:12: Fetch: https://github.com/mamemomonga.keys
	mkdir -m 0700 -p ~/.ssh
	cat >> ~/.ssh/authorized_keys << 'EOS'
	# -----------------------------------
	#  GitHub User: mamemomonga
	#  Created At: 2020-06-04 05:54:20
	ssh-ed25519 AAAAxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	ssh-ed25519 AAAAxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	# -----------------------------------
	
	EOS
	chmod 600 ~/.ssh/authorized_keys

MacOS: SSH公開鍵を設定するコマンドをクリップボードにペーストする

	$ fetch-github-ssh-key -u mamemomonga -g | pbcopy


