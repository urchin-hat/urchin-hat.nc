# urchin-hat.nc

A TCP server built with Go that streams animated urchin-hat ASCII art to your terminal. Just connect with `nc`!

`nc`（netcat）で接続すると、ターミナルに自分のキャラクター **urchin-hat** の文字アニメーションが流れるサーバーです。Go で実装し、Google Cloud の Compute Engine（GCE）で世界中から接続できるサービスとして公開する予定です。

## Status

現在はプロジェクトの準備段階です。サーバーの実装・デプロイ・接続先の公開はこれから行います。

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
