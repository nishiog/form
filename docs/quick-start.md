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

```json
{
  "api": {
    "post_url": "https://your-api.com/submit",
    "secret_key": "あなたのシークレットキー"
  }
}
```

## ✅ 完了！

`output/index.html` をブラウザで開いて使用できます。

詳細は [Windows利用者向けマニュアル](windows-manual.md) を参照してください。

