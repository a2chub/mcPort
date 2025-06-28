# 開発環境構築手順

MC-Portを開発するための基本的な環境構築手順を以下に示します。ここではUbuntu 22.04
を想定していますが、他のUnix系OSでも同様です。

1. **依存ツールのインストール**
   ```bash
   sudo apt-get update
   sudo apt-get install -y git golang make
   ```
2. **ソースコードの取得**
   ```bash
   git clone https://example.com/mcport.git
   cd mcport
   ```
3. **Goモジュールの取得**
   ```bash
   go mod download
   ```
4. **ビルドの実行**
   ```bash
   make build
   ```
   `bin/` 以下に `mcportd` と `mcportctl` が生成されます。
5. **設定ファイルの作成**
   `example/mcport.yaml` を参考に必要なパラメータを指定します。
6. **テストの実行**
   ```bash
   go test ./...
   ```

以上で開発に必要な最低限の準備が整います。詳細なビルドオプションや開発フローは
他のドキュメントを参照してください。
