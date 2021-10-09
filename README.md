### 説明
- クリーンアーキテクチャを勉強するために作成しました。

### ディレクトリ構成

```
app
    │
    ├── domain（エンティティ）
    │    └── 「データ構造体」
    │
    ├── infrastructure（フレームワーク/ドライバ）
    │    ├── db.go　・　・　・　「DBへの接続、クエリ実行、ORMの呼び出し」
    │    ├── config.go　・　・　・　「設定情報」
    │    └── router.go　・　・　・　「ルーティング」
    │
    ├── interface（インターフェイスアダプター）
    │    ├── database　・　・　・　「infra ↔︎ interfaceの境界線」
    │    └── controllers　・　・　・　「リクエストをUsecaseへ処理を依頼し、処理結果をもとにレスポンスを生成」
    │
    ├── usecase（ユースケース）
    │    ├── repository　・　・　・　「controller ↔︎ interactorの境界線」
    │    └── intaractor　・　・　・　「domainを適宜呼び出してユースケースを実現」
    │
    └── main.go　・　・　・　「サーバー起動」
```

### URL
  - https://repgram.com/

### 使用技術

- フロントエンド

  - React 17.0.1
    - TypeScript、Next.js

- バックエンド

  - Golang 1.17
    - Gin、GORM

- インフラ
  - Docker,docker-compose
  - AWS(IAM,VPC,ECS,ECR,RDS,Route53,ELB,S3,CloudWatch)
  - terraform
  - CircleCI

### 機能一覧

- ユーザー登録
- 写真投稿、編集、削除
- いいね登録、削除
- ページネーション
- 投稿検索機能
