
# 図書館サービス

本の貸し出しサービスです。gRPC apiで貸し出しリクエストを送信します。Restful APIを使用して、本の貸し出し状況を閲覧します。

## Quick start

実行例です。

```
% git clone https://github.com/Kynea0b/demo-protocol-buffers.git
% cd demo-protocol-buffers/sample-book-lending
```

```shell
% pwd
demo-protocol-buffers/sample-book-lending
```

サーバを起動します。

```shell
% make start_server
```

貸し出しリクエストを送ってみます。

```shell
% make sendborrow1
```

貸し出し状況を閲覧します。

```shell
% make get_lendinginfo
```
貸し出し状況を閲覧します。※こちらは仮実装です。

```shell
% make get_borrowedtime
```

図書館に新しい本を登録します。

```shell
% make register_book
```

新しく登録された本の貸し出しリクエストを送ります。

```shell
% make sendborrow2
```

## ブラウザ

コマンドラインからリクエスト送信する場合は、`curl`
下記のようにブラウザからhttpリクエストを送っても良き。
http://localhost:8090/hello/%22foobarbaz%22

## Protocol Documentation

[Protocol Documentation](https://kynea0b.github.io/demo-protocol-buffers/sample-book-lending/doc/index.html)


