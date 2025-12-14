# Power Automate連携ガイド

## 📋 概要

このフォームは **Microsoft Power Automate** のHTTPトリガー（Incoming Webhook）との連携を想定しています。

Power AutomateのWebhookは**CORSに対応している**ため、HTMLファイルを直接開いても送信できます。

## 🔧 Power Automate側の設定

### 1. 新しいフローを作成

1. **Power Automate** (https://make.powerautomate.com/) にサインイン
2. **「作成」** → **「自動化したクラウド フロー」** を選択
3. フロー名を入力（例: 「書類作成フォーム受信」）
4. スキップして空白のフローを作成

### 2. HTTPトリガーを追加

1. **「HTTP 要求の受信時」** トリガーを検索して追加
2. **「要求本文のJSONスキーマ」** に以下を設定：

```json
{
  "type": "object",
  "properties": {
    "secret": {
      "type": "string"
    },
    "case_id": {
      "type": ["integer", "null"]
    },
    "mode": {
      "type": "string"
    },
    "documents": {
      "type": "array",
      "items": {
        "type": "object",
        "properties": {
          "document_id": {
            "type": "string"
          },
          "document_name": {
            "type": "string"
          },
          "fields": {
            "type": "object"
          }
        }
      }
    }
  }
}
```

3. **保存** をクリック
4. **HTTP POST の URL** が表示される（例: `https://prod-XX.japaneast.logic.azure.com:443/workflows/...`）

### 3. アクションを追加

受信したデータをどう処理するか設定します。

#### 例1: SharePointリストに追加

```
新しいステップ → SharePoint → 項目の作成
- サイトのアドレス: 社内SharePointサイト
- リスト名: 書類管理リスト
- タイトル: @{triggerBody()?['documents'][0]['document_name']}
- 会社名: @{triggerBody()?['documents'][0]['fields']['company_name']}
```

#### 例2: Excelファイルに追記

```
新しいステップ → Excel Online (Business) → 表に行を追加
- 場所: OneDrive for Business
- ドキュメントライブラリ: OneDrive
- ファイル: /書類管理/受信データ.xlsx
- テーブル: 受信データテーブル
- 各列に動的コンテンツをマッピング
```

#### 例3: Teamsにメッセージを投稿

```
新しいステップ → Microsoft Teams → メッセージを投稿する
- チームID: 法務チーム
- チャネルID: 書類管理
- メッセージ: 
  新しい書類作成依頼が届きました
  案件ID: @{triggerBody()?['case_id']}
  書類数: @{length(triggerBody()?['documents'])}
```

#### 例4: Outlookメールを送信

```
新しいステップ → Outlook → メールの送信
- 宛先: legal@company.com
- 件名: 新規書類作成依頼 - @{triggerBody()?['documents'][0]['fields']['company_name']}
- 本文: 
  案件ID: @{triggerBody()?['case_id']}
  モード: @{triggerBody()?['mode']}
  
  書類一覧:
  @{join(body('Select')?['document_name'], ', ')}
```

### 4. シークレットキーの検証（推奨）

セキュリティのため、シークレットキーを検証します。

```
新しいステップ → 条件
- triggerBody()?['secret']
- 次の値に等しい
- YOUR_SECRET_KEY_HERE

「はいの場合」: 通常の処理
「いいえの場合」: エラー応答を返す
  → 応答 → 状態コード: 403 → 本文: {"error": "Invalid secret"}
```

## ⚙️ フォーム側の設定

### config.json の設定

```json
{
  "api": {
    "post_url": "https://prod-XX.japaneast.logic.azure.com:443/workflows/XXXXXXXX/triggers/manual/paths/invoke?api-version=2016-06-01&sp=%2Ftriggers%2Fmanual%2Frun&sv=1.0&sig=XXXXXXXXXX",
    "method": "POST",
    "headers": {
      "Content-Type": "application/json"
    },
    "secret_key": "your-secret-key-here"
  },
  "app": {
    "title": "書類作成フォーム",
    "autosave_interval_ms": 500,
    "enable_console_log": true
  }
}
```

**重要**: 
- `post_url` にはPower Automateで生成された完全なURL（クエリパラメータ含む）をコピー
- URLは非常に長いですが、すべてコピーしてください
- `secret_key` は任意の文字列（Power Automate側で検証する場合）

## 🎯 CORS問題について

### Power AutomateはCORS対応済み

Power Automate (Microsoft Flow) のHTTPトリガーは：
- ✅ **CORSに対応している**
- ✅ `file://` プロトコルからのリクエストも受け付ける
- ✅ 追加のCORS設定は不要

**つまり、HTMLファイルをダブルクリックで開いても送信できます！**

### 確認方法

1. `form-builder.exe` を実行してHTMLを生成
2. `output/index.html` をダブルクリックして開く
3. フォームに入力して「送信」ボタンをクリック
4. Power Automateのフロー実行履歴を確認

ローカルHTTPサーバーは**不要**です。

## 📊 送信されるデータ例

Power Automateが受信するJSONデータ：

```json
{
  "secret": "your-secret-key-here",
  "case_id": 1,
  "mode": "case",
  "documents": [
    {
      "document_id": "minutes_regular_meeting",
      "document_name": "定時株主総会議事録",
      "fields": {
        "company_name": "株式会社サンプル",
        "preparation_date": "2024-12-14",
        "head_office_address": "東京都千代田区丸の内一丁目1番1号",
        "representative_director_name": "山田太郎",
        "meeting_date": "2024-06-28",
        "meeting_time": "10:00",
        "meeting_place": "本社会議室",
        "agenda_items": "第1号議案...",
        "resolution_details": "承認された..."
      }
    }
  ]
}
```

## 🔄 Power Automateでのデータ処理例

### 動的コンテンツの参照方法

#### シークレットキー
```
@{triggerBody()?['secret']}
```

#### 案件ID
```
@{triggerBody()?['case_id']}
```

#### 最初の書類の名前
```
@{triggerBody()?['documents'][0]['document_name']}
```

#### 会社名（最初の書類から）
```
@{triggerBody()?['documents'][0]['fields']['company_name']}
```

#### すべての書類をループ処理

```
Apply to each
- 選択: @{triggerBody()?['documents']}
- アクション内で: @{items('Apply_to_each')?['document_name']}
```

### JSONデータの展開

複数の書類がある場合、各書類を個別に処理：

```
Apply to each: @{triggerBody()?['documents']}
  ├─ 書類ID: @{items('Apply_to_each')?['document_id']}
  ├─ 書類名: @{items('Apply_to_each')?['document_name']}
  └─ フィールド: @{items('Apply_to_each')?['fields']}
       ├─ 会社名: @{items('Apply_to_each')?['fields']?['company_name']}
       ├─ 作成日: @{items('Apply_to_each')?['fields']?['preparation_date']}
       └─ ...
```

## 🔒 セキュリティ対策

### 1. シークレットキーの検証

```
条件コントロール
IF @{triggerBody()?['secret']} equals "your-expected-secret"
  THEN: 通常処理
  ELSE: 
    応答アクション
    - 状態コード: 403
    - 本文: {"error": "Unauthorized"}
    - 終了
```

### 2. IPアドレス制限（オプション）

Power Automate Premiumプランの場合、特定のIPアドレスからのみアクセス許可可能。

### 3. URLの保護

- Power AutomateのWebhook URLには認証トークンが含まれている
- このURLを他人に知られないように管理
- 定期的にフローを再作成してURL更新を検討

## 📝 実運用の推奨フロー

### フロー例: 書類作成依頼の処理

```
1. HTTP 要求の受信時
   ↓
2. 条件: シークレットキー検証
   ↓（認証成功）
3. Apply to each: documents配列をループ
   ↓
4. SharePoint リストに項目を作成
   - 書類ID: @{items('Apply_to_each')?['document_id']}
   - 書類名: @{items('Apply_to_each')?['document_name']}
   - 会社名: @{items('Apply_to_each')?['fields']?['company_name']}
   - 作成日: @{items('Apply_to_each')?['fields']?['preparation_date']}
   - その他フィールド...
   ↓
5. Teams チャネルに通知
   ↓
6. 担当者にメール送信
   ↓
7. HTTP 応答
   - 状態コード: 200
   - 本文: {"status": "success", "message": "受信しました"}
```

## 🐛 トラブルシューティング

### エラー: 「ネットワークエラー」

**原因**: Webhook URLが間違っている

**解決方法**:
1. Power Automateのフローを開く
2. 「HTTP 要求の受信時」トリガーをクリック
3. URL全体をコピー（非常に長いです）
4. `config.json`の`post_url`に貼り付け
5. フォームを再生成（`form-builder.exe`を実行）

### エラー: 「403 Forbidden」

**原因**: シークレットキーが一致していない

**解決方法**:
1. `config.json`の`secret_key`を確認
2. Power Automateのフロー内の検証条件を確認
3. 値が完全に一致しているか確認（大文字小文字も区別）

### フローが実行されない

**確認事項**:
1. Power Automateのフローが**オン**になっているか
2. フローの実行履歴でエラーを確認
3. ブラウザの開発者ツール（F12）でネットワークエラーを確認

### データが正しく保存されない

**原因**: フィールド名のマッピングミス

**解決方法**:
1. Power Automateでテストデータを確認
2. フィールドIDを確認（`company_name`など）
3. 動的コンテンツのパスを修正

## 💡 応用例

### Office 365との統合

- **SharePoint**: ドキュメントライブラリに自動保存
- **OneDrive**: 個人用フォルダに整理
- **Teams**: チャネルに通知、承認ワークフロー
- **Outlook**: 関係者に自動メール送信
- **Planner**: タスクの自動作成

### データ処理の自動化

- **Word/Excel**: テンプレートに自動入力
- **PDF**: 自動生成と保存
- **承認フロー**: 上司の承認を経由
- **履歴管理**: すべての提出履歴をログ

## 📞 サポート

Power Automate関連の質問：
- [Microsoft Power Automate ドキュメント](https://docs.microsoft.com/ja-jp/power-automate/)
- [Power Automate コミュニティ](https://powerusers.microsoft.com/t5/Power-Automate-Community/ct-p/MPACommunity)

---

最終更新: 2024-12-14

