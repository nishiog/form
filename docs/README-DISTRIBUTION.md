# 書類作成フォームビルダー - 配布パッケージ

## 📦 このパッケージについて

このパッケージには、書類作成フォームを自動生成するために必要なすべてのファイルが含まれています。

**対象ユーザー**: Windows環境で使用する方（Go環境は不要）

## 📁 パッケージ内容

```
form-builder-windows/
├── form-builder.exe          # 実行ファイル
├── template.html             # フォームのテンプレート
├── README.txt                # 簡易説明
├── assets/                   # 設定・データファイル
│   ├── case.json             # 案件データ（69件）
│   ├── documents.json        # 書類データ（40種類）
│   ├── fields.json           # 入力項目（65個）
│   ├── case_documents.json   # 案件と書類の紐づけ
│   ├── document_fields.json  # 書類とフィールドの紐づけ
│   └── config.json.example   # 設定ファイルのサンプル
├── docs/                     # マニュアル
│   ├── windows-manual.md     # 詳細マニュアル
│   └── quick-start.md        # クイックスタート
└── output/                   # 生成されたHTMLの出力先（初期は空）
```

## 🎯 機能

### 案件選択モード
- 69種類の案件から選択
- 必要な書類が自動で選択される
- 重複する入力項目は1回だけ入力

### 手動選択モード
- 40種類の書類から個別に選択
- 必要な書類だけを選んで使用

### 自動生成される内容
- 入力フォーム（HTML形式）
- 案件に応じた入力項目
- API送信機能
- LocalStorage自動保存

## 🚀 使用方法

### 初回のみ: 設定ファイルの作成

1. `assets/config.json.example` を `assets/config.json` にコピー
2. メモ帳で開いて以下を編集：

#### Power Automateを使用する場合（推奨）

```json
{
  "api": {
    "post_url": "https://prod-XX.japaneast.logic.azure.com:443/workflows/.../triggers/manual/paths/invoke?api-version=2016-06-01&sp=%2Ftriggers%2Fmanual%2Frun&sv=1.0&sig=XXXXX",
    "secret_key": "your-secret-key-here"
  }
}
```

**メリット**:
- ✅ HTMLファイルを直接開いても送信できる（CORSの心配なし）
- ✅ SharePoint/Teams/Outlookと簡単に連携
- ✅ 追加のサーバー構築不要

詳細: `docs/power-automate-setup.md`

#### その他のAPIサーバーを使用する場合

```json
{
  "api": {
    "post_url": "https://your-api-server.com/submit",
    "secret_key": "your-secret-key-here"
  }
}
```

注意: CORS設定が必要な場合があります。
### フォームの生成

1. `form-builder.exe` をダブルクリック
2. `output/index.html` が生成されます
3. ブラウザで開いて使用

## 📝 送信されるデータ形式

```json
{
  "secret": "your-secret-key",
  "case_id": 1,
  "mode": "case",
  "documents": [
    {
      "document_id": "minutes_regular_meeting",
      "document_name": "定時株主総会議事録",
      "fields": {
        "company_name": "株式会社サンプル",
        "preparation_date": "2024-12-14",
        "meeting_date": "2024-06-28"
      }
    }
  ]
}
```

## 🔐 セキュリティ

### 重要な注意事項

⚠️ `config.json` には機密情報が含まれます
- このファイルを他人と共有しないでください
- 定期的にシークレットキーを変更してください

⚠️ 生成された `output/index.html` にもシークレットキーが含まれます
- 不特定多数がアクセスできる場所には置かないでください
- 配布する場合は適切なアクセス制限を設定してください

## 📚 マニュアル

- **詳細マニュアル**: `docs/windows-manual.md`
- **クイックスタート**: `docs/quick-start.md`

## 💡 ヒント

### データのカスタマイズ

`assets` フォルダ内のJSONファイルを編集することで、案件や書類をカスタマイズできます。

編集後は `form-builder.exe` を実行して再生成してください。

### バックアップ

定期的に `assets` フォルダをバックアップすることをお勧めします。

### アップデート

新しいバージョンが公開されたら：
1. 新しい `form-builder-windows.zip` をダウンロード
2. 現在の `assets/config.json` を新しいフォルダにコピー
3. 必要に応じて他の `assets/*.json` もコピー

## 🐛 トラブルシューティング

### よくあるエラー

**「設定ファイルの読み込みエラー」**
→ `assets/config.json` が存在しません。初期設定を行ってください。

**「JSONのパースエラー」**
→ JSONファイルの形式が正しくありません。カンマや括弧を確認してください。

**「送信エラー」**
→ `config.json` のAPI URLとシークレットキーを確認してください。

詳細は `docs/windows-manual.md` の「トラブルシューティング」を参照してください。

## 📞 サポート

問題が解決しない場合は、GitHubのIssuesで質問してください。

## 📄 ライセンス

このソフトウェアのライセンスについては、開発者にお問い合わせください。

---