# ドキュメント

書類作成フォームビルダーのドキュメント一覧です。

## 📚 利用者向けマニュアル

### Windows利用者向け

- **[Windows利用者向けマニュアル](windows-manual.md)** ⭐ 必読
  - インストール手順
  - 初期設定方法
  - 使い方
  - トラブルシューティング
  - 詳細な解説

- **[Power Automate連携ガイド](power-automate-setup.md)** 🔥 推奨
  - Microsoft Power Automateとの連携手順
  - Webhook URLの取得方法
  - SharePoint/Teams/Outlookとの統合例
  - CORSの心配不要（HTMLを直接開いてOK）

- **[クイックスタートガイド](quick-start.md)**
  - 3ステップで始める
  - 最小限の設定で使用開始

- **[配布パッケージ説明書](README-DISTRIBUTION.md)**
  - パッケージの内容
  - 機能の概要
  - データ形式の説明

## 🔧 開発者向けドキュメント

- **[リリースプロセス](release-process.md)**
  - GitHub Actionsによる自動ビルド
  - バージョン管理
  - リリース手順
  - トラブルシューティング

- **[開発者向けREADME](../README.md)**
  - プロジェクト概要
  - ファイル構成
  - データ構造
  - カスタマイズ方法

## 📖 ドキュメント構成

```
docs/
├── README.md                    # このファイル
├── windows-manual.md            # Windows利用者向け詳細マニュアル
├── quick-start.md               # クイックスタート
├── README-DISTRIBUTION.md       # 配布パッケージ説明
└── release-process.md           # リリースプロセス（開発者向け）
```

## 🎯 対象読者

### Windows利用者（Go環境なし）
- [Windows利用者向けマニュアル](windows-manual.md) から始めてください
- Go言語の知識は不要です
- 実行ファイル（.exe）を提供しています

### 開発者（Go環境あり）
- [開発者向けREADME](../README.md) を参照してください
- JSONファイルのカスタマイズ
- テンプレートの編集
- リリース作成

## 🚀 はじめに

### Windows利用者の方

1. GitHubのReleasesから `form-builder-windows.zip` をダウンロード
2. [クイックスタートガイド](quick-start.md) に従って設定
3. 詳細は [Windows利用者向けマニュアル](windows-manual.md) を参照

### 開発者の方

1. リポジトリをクローン
2. [開発者向けREADME](../README.md) に従って環境構築
3. リリースは [リリースプロセス](release-process.md) を参照

## 💡 よくある質問

**Q: どのマニュアルを読めばいいですか？**
- Windows利用者: [Windows利用者向けマニュアル](windows-manual.md)
- 急いでいる方: [クイックスタートガイド](quick-start.md)
- 開発者: [開発者向けREADME](../README.md)

**Q: Goをインストールする必要がありますか？**
- Windows利用者: 不要です。実行ファイルを提供しています
- 開発者: 必要です（Go 1.21以降）

**Q: パッケージをカスタマイズできますか？**
- はい。assetsフォルダのJSONファイルを編集できます
- 詳細は各マニュアルの「データのカスタマイズ」を参照

## 📞 サポート

問題が解決しない場合：
1. 該当するマニュアルの「トラブルシューティング」を確認
2. GitHubのIssuesで質問を投稿

---