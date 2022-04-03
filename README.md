# grpc-graphql-server

このリポジトリはGraphQLとgRPCをフル活用したマイクロサービスアーキテクチャの練習用リポジトリです。細々と機能実装をしてきます。

## 作ろうとしているもの

簡単なチャットサーバーを実装します。

## テスト

curl --request POST \
    --header 'content-type: application/json' \
    --url http://localhost:8082/query \
    --data '{"query":"query{\n  users {\n    id\n    name\n  }\n}"}'