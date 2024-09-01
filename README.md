# 非公式Misskey CLIツール

## これなに
`tootctl`、`fishctl`にインスパイアを受けた、Misskey/CherryPickで使える非公式Go言語製CLIツールです。

## インストール

### 必要なもの
- Go 1.23

```bash
git clone https://github.com/1673beta/go-cmd-prac.git
cd go-cmd-prac
go build
export PATH=$PATH:/this/repository/path
```

## 使い方
 
 ### アップデートする
```bash
mkctl update [repository] [branch] [-l]
```
repositoryにリポジトリURLを指定することができます。指定しない場合、https://github.com/misskey-dev/misskey.git からpullされます。

branchには、pullするブランチを指定することができます。デフォルトでは`master`です。

`-l`は、低メモリ環境で実行する際のオプションです。3072MBのメモリ領域を確保して実行します。
