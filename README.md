# conoha-cli
ConoHa VPS 管理API CLIツール

## 機能
- アカウント管理
    - 入金残高確認
    - 入金履歴確認
    - 利用サービス確認
    - 請求履歴確認
- VM管理
    - VM確認
    - VMプラン(Flavor)確認
- セキュリティグループ管理
    - セキュリティグループ確認
    - セキュリティグループルール確認

## 動作環境
- Go 1.14

## 使用方法
### ビルド
```shell
git clone https://github.com/aqmr-kino/conoha-cli.git
```

```shell
# for Mac/Linux
go build -o conoha-cli main.go

# for Windows
go build -o conoha-cli.exe main.go
```

### 初期設定
#### ConoHa APIアカウント設定
```shell
conoha-cli config --user username --password password --tenant tenand-id --endpoint https://identity.hogehoge.com/v2.0
```

#### APIエンドポイント設定
```shell
# Account
conoha-cli set-endpoint --account https://account.hogehoge.com/v1/0123456789abcdef

# Compute
conoha-cli set-endpoint --compute https://compute.hogehoge.com/v2/0123456789abcdef

# Network
conoha-cli set-endpoint --network https://network.hogehoge.com
```

### アカウント関連

#### 入金残高確認
```shell
conoha-cli billing get-deposit
```

#### 入金履歴確認
```shell
conoha-cli billing list-deposit-history
```

#### 利用サービス確認
```shell
conoha-cli billing list-item
```

#### 請求履歴確認
```shell
conoha-cli billing list-invoice
```

### VM関連

#### VM確認
```shell
# 簡易(名前+状態のみ)
conoha-cli vm list

# 詳細表示
conoha-cli vm list --detail
```

#### VMプラン(Flavor)確認
```shell
conoha-cli vm list-flavor
```

### セキュリティグループ関連

#### 一覧取得
```shell
# 簡易(名前+IDのみ)
conoha-cli secgroup list

# 詳細表示
conoha-cli secgroup list --detail
```

#### ルール一覧取得
```shell
# 簡易(概要のみ)
conoha-cli secgroup list-rule

# 詳細表示
conoha-cli secgroup list-rule --detail
```

## 設定ファイル等
|ファイル名|内容|
|:--|:--|
|$HOME/.conoha|アカウント情報、エンドポイント情報|

## 変更履歴
### v0.3.0 (2020-05-17)
- 新機能
    - VM管理機能を追加

### v0.2.0 (2020-05-06)
- 新機能
    - アカウント管理機能を追加
- 修正
    - 一部サブコマンドのエラー処理不備を修正
    - Identity APIエンドポイントにバージョン番号まで含めるよう修正

### v0.1.0 (2020-05-03)
- 初版
