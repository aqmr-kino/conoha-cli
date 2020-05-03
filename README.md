# conoha-cli
ConoHa VPS 管理API CLIツール

## 機能
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
conoha-cli config --user username --password password --tenant tenand-id --endpoint https://identity.hogehoge.com
```

#### APIエンドポイント設定
```shell
conoha-cli set-endpoint --network https://network.hogehoge.com
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
### v0.1.0 (2020-05-03)
- 初版
