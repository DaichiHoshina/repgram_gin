### 説明
- クリーンアーキテクチャを勉強するために作成中。

### ディレクトリ構成
└── src
    ├── app
    │   ├── domain
    │   ├── infrastructure
    │   ├── interfaces
    │   │   ├── controllers
    │   │   └── database
    │   ├── main.go
    │   └── usecase

### URL
  - https://repgram.com/(停止中)

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

- ユーザー登録、編集
- 写真投稿、編集、削除
- いいね登録、削除
