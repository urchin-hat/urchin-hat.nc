# urchin-hat.nc

A planned TCP server built with Go that streams animated urchin-hat ASCII art to your terminal via `nc`.

`nc`（netcat）で接続すると、ターミナルに自分のキャラクター **urchin-hat** の文字アニメーションが流れるサーバーです。Go で実装し、Google Cloud の Compute Engine（GCE）で世界中から接続できるサービスとして公開する予定です。

## Status

mise による Go 開発環境と Go モジュールを初期化しています。サーバーの実装・デプロイ・接続先の公開はこれから行います。

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

Go のソースコードを追加した後は、次のコマンドで整形・テスト・ビルドを実行できます。

```sh
mise exec -- go fmt ./...
mise exec -- go test ./...
mise exec -- go build ./...
```

現時点では Go のパッケージやテストはありません。`go test ./...` は対象パッケージがないため失敗します。サーバーの起動コマンドは実装時に追記します。

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

- Go による TCP サーバー
- urchin-hat の ASCII アニメーション配信
- 複数クライアントの同時接続
- Google Cloud Compute Engine へのデプロイ

## Blog

制作過程は [urchin-hat のテックブログ](https://urchin-hat.org/memo/)で紹介する予定です。

記事タイトル（予定）：

> Google Cloud (GCE)で作る！世界中から nc で接続できる文字アニメーション配信サーバー
