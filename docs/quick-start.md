# クイックスタートガイド

## 🚀 3ステップで始める

### ステップ1: ダウンロード

GitHubのReleasesページから `form-builder-windows.zip` をダウンロードして解凍

### ステップ2: 設定

```
1. assets/config.json.example をコピー
2. config.json にリネーム
3. メモ帳で開いて編集
   - post_url: API URL
   - secret_key: シークレットキー
```

### ステップ3: 実行

```
form-builder.exe をダブルクリック
↓
output/index.html が生成される
↓
ブラウザで開く
```

## 📋 必要な設定（config.json）

### Power Automateを使用する場合（推奨）

```json
{
  "api": {
    "post_url": "https://prod-XX.japaneast.logic.azure.com:443/workflows/.../triggers/manual/paths/invoke?...",
    "secret_key": "あなたのシークレットキー"
  }
}
```

**メリット**: CORSの心配なし！HTMLをダブルクリックで開いても送信できます。

詳細: [Power Automate連携ガイド](power-automate-setup.md)

### その他のAPIを使用する場合

```json
{
  "api": {
    "post_url": "https://your-api.com/submit",
    "secret_key": "あなたのシークレットキー"
  }
}
```

注意: CORS設定が必要な場合があります。

## ✅ 完了！

`output/index.html` をブラウザで開いて使用できます。

詳細は [Windows利用者向けマニュアル](windows-manual.md) を参照してください。

