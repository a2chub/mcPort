# パッケージ作成手順書

MC-Portを配布するためのビルドおよびパッケージ作成の流れを示します。ここでは
Linux環境でのdxtパッケージ生成を例に説明します。

1. **ビルド準備**
   - 必要なコンパイラとライブラリをインストール (`go`, `make` 等)。
   - ソースコードをクリーン状態にするため `git clean -fdx` を実行。
2. **バイナリのビルド**
   - `make build` で `mcportd` と `mcportctl` を生成。クロスビルドが必要な場合は
     `GOOS` `GOARCH` などの環境変数を指定。
3. **ファイル配置**
   - `bin/` 以下にビルド成果物を配置し、設定テンプレートやドキュメントを同梱する。
4. **dxt パッケージ生成**
   - `tools/dxt_init.sh` を実行して `dxt.yaml` を生成し、パッケージ名やバージョン、依存関係を
     記述する。以下は最小例です。

```yaml
name: mc-port
version: 1.0.0
arch: amd64
summary: dynamic port allocator
```

   - メタデータを編集後、`tools/dxt_build.sh` を実行してアーカイブを生成する。
   - 生成されたファイルは `dist/` ディレクトリへ出力。
5. **配布**
   - リリースノートを作成し、GitHub Releases などにアップロード。
   - 必要に応じてパッケージ署名を行う。

パッケージング用スクリプト（`dxt_init.sh`, `dxt_build.sh`）は `tools/` ディレクトリに配置しています。
