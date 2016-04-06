[![wercker status](https://app.wercker.com/status/52243183472f21dc72756d12a649ee19/m/master "wercker status")](https://app.wercker.com/project/bykey/52243183472f21dc72756d12a649ee19)

# soracom-cli

SORACOM API を呼び出すためのコマンドラインツール soracom を提供します。

# 特徴

soracom コマンドは以下のような特徴を備えています。

- soracom コマンドのバイナリファイルは API 定義ファイルから自動生成されますので、新しい API がリリースされた場合も迅速に対応が可能です。

- Go でクロスコンパイルされたバイナリファイルをターゲットの環境にコピーするだけで実行できます。環境を構築したり依存関係を解決したりする必要がありません。

- 指定された引数を元にリクエストを組み立て、SORACOM API を呼び出します。API からのレスポンス (JSON) をそのまま標準出力へ出力します。
  - これにより、soracom コマンドの出力を加工して他のコマンドへ渡したりすることが容易にできるようになります。

- bash completion（引数補完）に対応しています。以下のような行を .bashrc 等に記述してください。
  ```
  eval "$(soracom completion)"
  ```


# インストール方法

## Go を必要としない方法

[Releases のページ](https://github.com/soracom/soracom-cli/releases) からターゲットの環境に合ったパッケージファイルをダウンロードして展開し、実行形式ファイルを PATH の通ったディレクトリに配置します。


## Go を使う方法

```
go get -u github.com/soracom/soracom-cli/soracom
```


# ビルド/テスト方法

ソースからビルドしたい開発者の方や、バグ修正/機能追加等の Pull Request をしたい場合は以下のいずれかの方法でビルドおよびテストを行ってください。

## ローカル環境でビルドする方法 (Linux / Mac OS X)

Go がインストールされている状態で、以下のようにビルドスクリプトを実行します。

```
./scripts/build.sh 1.2.3
```

ここで 1.2.3 はバージョン番号です。適当な番号を指定してください。

ビルドが成功したら、次にテストを実行します。

```
./test/test.sh
```


## wercker を使ってビルドする方法

wercker の CLI をインストールし、以下のようにビルドを実行します。テストまで自動的に実行されます。

```
wercker build
```

TODO: 現状、ビルド結果はコンテナの中に出力されるので、マウントしたボリュームに出力できるように修正予定です
