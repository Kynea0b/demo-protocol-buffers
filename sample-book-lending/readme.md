
# 図書館サービス

本の貸し出しサービスです。gRPC apiで貸し出しリクエストを送信します。Restful APIを使用して、本の貸し出し状況を閲覧します。

## Quick start

実行例です。

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
% make call_sendborrow_method_of_the_service
```

貸し出し状況を閲覧します。※こちらは仮実装です。

```shell
% make show_accountinfo
```

## document

[document](https://kynea0b.github.io/demo-protocol-buffers/sample-book-lending/doc/index.html)


