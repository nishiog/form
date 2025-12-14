# 書類作成フォームビルダー - Windows利用者向けマニュアル

## 📋 概要

このツールは、案件に応じた書類作成フォームをHTML形式で自動生成するアプリケーションです。
Windows環境で簡単に使用できます。

## 💻 必要な環境

- Windows 10以降
- インターネット接続（API送信を行う場合）
- Webブラウザ（Chrome、Edge、Firefoxなど）

**注意**: Goのインストールは不要です。実行ファイル（.exe）を提供しています。

## 📦 インストール手順

### 1. パッケージのダウンロード

GitHubのReleasesページから最新版の `form-builder-windows.zip` をダウンロードします。

```
https://github.com/[あなたのユーザー名]/[リポジトリ名]/releases
```

### 2. ファイルの解凍

ダウンロードしたZIPファイルを右クリック → 「すべて展開」で解凍します。

推奨フォルダ: `C:\form-builder\`

### 3. フォルダ構成の確認

解凍後、以下のファイルとフォルダが含まれていることを確認してください：

```
form-builder/
├── form-builder.exe          # 実行ファイル
├── template.html             # HTMLテンプレート
├── README.txt                # 簡易マニュアル
├── assets/                   # 設定・データファイル
│   ├── case.json
│   ├── documents.json
│   ├── fields.json
│   ├── case_documents.json
│   ├── document_fields.json
│   └── config.json.example   # 設定ファイルのサンプル
├── docs/                     # ドキュメント
│   └── windows-manual.md     # このファイル
└── output/                   # 生成されたHTMLの出力先（空）
```

## ⚙️ 初期設定

### 設定ファイルの作成

**重要**: 初回のみ、設定ファイルを作成する必要があります。

1. **エクスプローラーで `assets` フォルダを開く**

2. **`config.json.example` を見つける**

3. **ファイルをコピー**
   - `config.json.example` を右クリック → 「コピー」
   - 同じフォルダ内で右クリック → 「貼り付け」
   - コピーされたファイルの名前を `config.json` に変更

4. **`config.json` を編集**
   - `config.json` を右クリック → 「プログラムから開く」 → 「メモ帳」
   - 以下の項目を編集します：

```json
{
  "api": {
    "post_url": "https://your-api-server.com/submit",
    "method": "POST",
    "headers": {
      "Content-Type": "application/json"
    },
    "secret_key": "ここに実際のシークレットキーを入力"
  },
  "app": {
    "title": "書類作成フォーム",
    "autosave_interval_ms": 500,
    "enable_console_log": true
  }
}
```

### 設定項目の説明

| 項目 | 説明 | 例 |
|------|------|-----|
| `post_url` | 送信先のAPI URL | `https://api.example.com/documents` |
| `method` | HTTPメソッド | `POST` |
| `secret_key` | API認証用のキー | システム管理者から提供されたキー |
| `enable_console_log` | デバッグログの有効化 | `true` または `false` |

## 🚀 使い方

### フォームの生成

#### 方法1: ダブルクリックで実行（簡単）

1. `form-builder.exe` をダブルクリック
2. コマンドプロンプトのウィンドウが開きます
3. 処理が完了すると以下のメッセージが表示されます：

```
✅ HTMLファイルが生成されました: output\index.html
   案件数: 69
   書類数: 40
   フィールド数: 65
```

4. 何かキーを押すとウィンドウが閉じます

#### 方法2: コマンドプロンプトから実行（詳細確認）

1. **スタートメニュー** → 「cmd」と入力 → **コマンドプロンプト**を起動

2. **フォルダに移動**
```cmd
cd C:\form-builder
```

3. **実行**
```cmd
form-builder.exe
```

4. **結果の確認**
```
✅ HTMLファイルが生成されました: output\index.html
   案件数: 69
   書類数: 40
   フィールド数: 65
```

### 生成されたフォームの使用

1. **`output` フォルダを開く**

2. **`index.html` をブラウザで開く**
   - `index.html` をダブルクリック
   - または、右クリック → 「プログラムから開く」 → お好みのブラウザ

3. **フォームを使用**
   - 案件を選択
   - 必要な情報を入力
   - 「送信」ボタンをクリック

## 📝 データのカスタマイズ

### 案件データの編集

`assets` フォルダ内のJSONファイルを編集することで、案件や書類をカスタマイズできます。

| ファイル | 内容 | 編集推奨度 |
|----------|------|-----------|
| `case.json` | 案件のリスト | ⭐⭐⭐ |
| `documents.json` | 書類の定義 | ⭐⭐ |
| `fields.json` | 入力項目の定義 | ⭐⭐ |
| `case_documents.json` | 案件と書類の紐づけ | ⭐⭐⭐ |
| `document_fields.json` | 書類と入力項目の紐づけ | ⭐⭐ |

**編集方法**:
1. JSONファイルを「メモ帳」または「Visual Studio Code」などで開く
2. JSON形式を崩さないように注意して編集
3. 保存後、`form-builder.exe` を再実行

**注意**: JSON形式のエラーがあるとビルドに失敗します。

## 🔍 トラブルシューティング

### エラー: 「設定ファイルの読み込みエラー」

**原因**: `assets/config.json` が存在しない

**解決方法**:
1. `assets/config.json.example` を `assets/config.json` にコピー
2. 「初期設定」の手順に従って設定

### エラー: 「JSONのパースエラー」

**原因**: JSONファイルの形式が正しくない

**解決方法**:
1. エラーメッセージでどのファイルが問題か確認
2. オンラインのJSON検証ツールでチェック（例: https://jsonlint.com/）
3. カンマの位置、引用符、括弧の閉じ忘れなどを確認

### フォームが生成されない

**確認事項**:
- [ ] `form-builder.exe` がある場所で実行しているか
- [ ] `template.html` が同じフォルダにあるか
- [ ] `assets` フォルダが存在するか
- [ ] `assets/config.json` が存在するか

### 送信エラーが発生する

**確認事項**:
- [ ] `config.json` の `post_url` が正しいか
- [ ] `secret_key` が正しく設定されているか
- [ ] インターネット接続があるか
- [ ] APIサーバーが稼働しているか

## 🔐 セキュリティに関する注意

### 重要: config.jsonの取り扱い

- `config.json` には機密情報（`secret_key`）が含まれます
- **このファイルを他人と共有しないでください**
- 定期的にシークレットキーを変更してください

### 生成されたHTMLファイルについて

- `output/index.html` には `secret_key` が埋め込まれています
- 配布する際は、アクセス制限を設定してください
- 不特定多数がアクセスできる場所には置かないでください

## 📞 サポート

### よくある質問

**Q: 複数のパソコンで使用できますか？**
A: はい。ただし、各パソコンで初期設定（config.jsonの作成）が必要です。

**Q: データをバックアップしたい**
A: `assets` フォルダ全体をコピーしてバックアップしてください。

**Q: 新しいバージョンにアップデートしたい**
A: GitHubから最新版をダウンロードし、古い `assets/config.json` を新しいフォルダにコピーしてください。

**Q: Mac/Linuxでも使えますか？**
A: このパッケージはWindows専用です。Mac/Linuxをお使いの場合は、Go環境をインストールして開発者向けの方法で使用してください。

### 問題が解決しない場合

GitHubのIssuesページで質問してください：
```
https://github.com/[ユーザー名]/[リポジトリ名]/issues
```

## 📚 参考資料

- [開発者向けREADME](../README.md)
- [GitHub Release](https://github.com/[ユーザー名]/[リポジトリ名]/releases)

---