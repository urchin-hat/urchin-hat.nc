# urchin-hat.nc

A planned TCP server built with Go that streams animated urchin-hat ASCII art to your terminal via `nc`.

`nc`（netcat）で接続すると、ターミナルに自分のキャラクター **urchin-hat** の文字アニメーションが流れるサーバーです。Go で実装し、Google Cloud の Compute Engine（GCE）で世界中から接続できるサービスとして公開する予定です。

## Status

TCP 接続に `hello,world` を返す最小構成のサーバーを実装しています。ASCII アニメーションの配信・デプロイ・接続先の公開はこれから行います。

## Development

### セットアップ

[mise の公式手順](https://mise.jdx.dev/getting-started.html)に従って mise をインストールしてから、次を実行します。

```sh
git clone https://github.com/urchin-hat/urchin-hat.nc.git
cd urchin-hat.nc
mise trust
mise install
mise exec -- go version
```

Go のバージョンは `mise.toml` で **1.27.3** に固定しています。Go モジュール名は `github.com/urchin-hat/urchin-hat.nc` です。個人用の設定は Git 管理対象外の `mise.local.toml` に記載できます。

`mise exec --` を使うと、シェルの設定に依存せずプロジェクトで指定した Go を実行できます。通常の `go` コマンドとして使いたい場合は、[シェルの有効化手順](https://mise.jdx.dev/getting-started.html)を参照してください。

### 開発時のコマンド

次のコマンドで整形・テスト・ビルドを実行できます。

```sh
mise exec -- go fmt ./...
mise exec -- go test ./...
mise exec -- go build -o /tmp/urchin-hat-nc .
```

### サーバーの起動方法

セットアップ後、リポジトリのルートで実行します。

```sh
mise exec -- go run .
```

次のログが表示されたら、接続を受け付けています（先頭に日時が付きます）。

```text
listening on 127.0.0.1:2323
```

このターミナルは起動したままにします。停止するには `Ctrl+C` を押します。

### nc での確認方法

同じマシンの別のターミナルで接続します。`nc`（netcat）がインストールされている必要があります。

```sh
nc 127.0.0.1 2323
```

次の文字列が表示され、`nc` が終了してプロンプトに戻れば正常です。

```text
hello,world
```

サーバーは `hello,world` と改行を送信して接続を閉じ、次の接続を受け付けます。同じ `nc` コマンドを再実行して、繰り返し接続できることも確認できます。

`Connection refused` が表示される場合は、サーバーが起動していることと、起動ログのアドレス・ポートが接続先と一致していることを確認してください。

### 待受アドレス・ポートを変更する場合

デフォルトの待受アドレスは `127.0.0.1:2323` です。変更する場合は `-addr` を指定します。

```sh
mise exec -- go run . -addr 127.0.0.1:12345
```

この場合、確認用の接続先も変更します。

```sh
nc 127.0.0.1 12345
```

### 自動テストでの確認

```sh
mise exec -- go test -race ./...
```

TCP 接続時の応答内容、応答後の切断、繰り返し接続できることを検証します。

## How it works

```text
Your terminal                    Google Cloud / GCE
┌──────────────────┐   TCP      ┌──────────────────────┐
│ nc <HOST> <PORT>  │ ────────► │ Go TCP server        │
│                  │ ◄──────── │ urchin-hat animation │
└──────────────────┘  ASCII art └──────────────────────┘
```

専用クライアントは不要です。TCP 接続を通じて文字アニメーションを受け取り、ターミナルで表示します。

## Connect (planned)

接続先の公開後、`nc` が利用できるターミナルから次の形式で接続できます。

```sh
nc <HOST> <PORT>
```

`<HOST>` と `<PORT>` は公開時に案内するホスト名とポート番号に置き換えてください。現時点では接続できません。

終了するには `Ctrl+C` を押します。

## Planned features

- urchin-hat の ASCII アニメーション配信
- Google Cloud Compute Engine へのデプロイ

## Blog

制作過程は [urchin-hat のテックブログ](https://urchin-hat.org/memo/)で紹介する予定です。

記事タイトル（予定）：

> Google Cloud (GCE)で作る！世界中から nc で接続できる文字アニメーション配信サーバー
