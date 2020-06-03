# fetch-github-ssh-key

* GitHubの任意ユーザのSSH公開鍵を取得し、ペーストで .ssh/authorized_keys に設定するツール
* -g オプションをつけると、authorized_keysに追記するスクリプトを生成します。

## 使い方

	$ fetch-github-ssh-key

SSH公開鍵を表示する

	$ fetch-github-ssh-key -u mamemomonga

SSH公開鍵を設定するコマンドを表示する

	$ fetch-github-ssh-key -u mamemomonga -g

MacOS: SSH公開鍵を設定するコマンドをクリップボードにペーストする

	$ fetch-github-ssh-key -u mamemomonga -g | pbcopy


