# 構成
- `sqs/sqs/main.go`にLambdaで実際に動くコードが記載されている
- デプロイは[Serverless Flamework](https://www.serverless.com/)を使用している
  - `make deploy`でビルドしてデプロイされる
  - ローカルで実行する場合Serverless Flameworkのインストールが必要
