package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
)

// 案件のデータ構造
type Case struct {
	ID       int      `json:"id"`
	Category string   `json:"category"`
	Path     []string `json:"path"`
	Name     string   `json:"name"`
}

// 書類のデータ構造
type Document struct {
	ID     string `json:"id"`
	NameJA string `json:"name_ja"`
	NameEN string `json:"name_en"`
}

// フィールドのデータ構造
type Field struct {
	ID       string `json:"id"`
	LabelJA  string `json:"label_ja"`
	LabelEN  string `json:"label_en"`
	Type     string `json:"type"`
	Required bool   `json:"required"`
	Sample   string `json:"sample"`
	Priority int    `json:"priority"`
}

// 案件と書類の紐づけ
type CaseDocument struct {
	CaseID      int      `json:"case_id"`
	DocumentIDs []string `json:"document_ids"`
}

// 書類とフィールドの紐づけ
type DocumentField struct {
	DocumentID string   `json:"document_id"`
	FieldIDs   []string `json:"field_ids"`
}

// 設定ファイルのデータ構造
type Config struct {
	API struct {
		PostURL   string            `json:"post_url"`
		Method    string            `json:"method"`
		Headers   map[string]string `json:"headers"`
		SecretKey string            `json:"secret_key"`
	} `json:"api"`
	App struct {
		Title              string `json:"title"`
		AutosaveIntervalMs int    `json:"autosave_interval_ms"`
		EnableConsoleLog   bool   `json:"enable_console_log"`
	} `json:"app"`
}

// テンプレート用のデータ構造
type TemplateData struct {
	CaseDataJSON         template.JS
	DocumentsJSON        template.JS
	CaseDocumentsJSON    template.JS
	DocumentFieldsJSON   template.JS
	FieldsJSON           template.JS
	ConfigJSON           template.JS
}

func main() {
	// JSONファイルの読み込み
	cases, err := loadCases("assets/case.json")
	if err != nil {
		fmt.Printf("案件データの読み込みエラー: %v\n", err)
		os.Exit(1)
	}

	documents, err := loadDocuments("assets/documents.json")
	if err != nil {
		fmt.Printf("書類データの読み込みエラー: %v\n", err)
		os.Exit(1)
	}

	fields, err := loadFields("assets/fields.json")
	if err != nil {
		fmt.Printf("フィールドデータの読み込みエラー: %v\n", err)
		os.Exit(1)
	}

	caseDocuments, err := loadCaseDocuments("assets/case_documents.json")
	if err != nil {
		fmt.Printf("案件-書類紐づけデータの読み込みエラー: %v\n", err)
		os.Exit(1)
	}

	documentFields, err := loadDocumentFields("assets/document_fields.json")
	if err != nil {
		fmt.Printf("書類-フィールド紐づけデータの読み込みエラー: %v\n", err)
		os.Exit(1)
	}

	config, err := loadConfig("assets/config.json")
	if err != nil {
		fmt.Printf("設定ファイルの読み込みエラー: %v\n", err)
		fmt.Println("\n初回セットアップの場合:")
		fmt.Println("  cp assets/config.json.example assets/config.json")
		fmt.Println("  vim assets/config.json")
		os.Exit(1)
	}

	// 案件データをマップに変換
	caseMap := make(map[int]Case)
	for _, c := range cases {
		caseMap[c.ID] = c
	}

	// 書類データをマップに変換
	documentMap := make(map[string]Document)
	for _, d := range documents {
		documentMap[d.ID] = d
	}

	// フィールドデータをマップに変換
	fieldMap := make(map[string]Field)
	for _, f := range fields {
		fieldMap[f.ID] = f
	}

	// JSONに変換
	caseDataJSON, err := toJSON(caseMap)
	if err != nil {
		fmt.Printf("案件データのJSON変換エラー: %v\n", err)
		os.Exit(1)
	}

	documentsJSON, err := toJSON(documentMap)
	if err != nil {
		fmt.Printf("書類データのJSON変換エラー: %v\n", err)
		os.Exit(1)
	}

	fieldsJSON, err := toJSON(fieldMap)
	if err != nil {
		fmt.Printf("フィールドデータのJSON変換エラー: %v\n", err)
		os.Exit(1)
	}

	caseDocumentsJSON, err := toJSON(caseDocuments)
	if err != nil {
		fmt.Printf("案件-書類紐づけデータのJSON変換エラー: %v\n", err)
		os.Exit(1)
	}

	documentFieldsJSON, err := toJSON(documentFields)
	if err != nil {
		fmt.Printf("書類-フィールド紐づけデータのJSON変換エラー: %v\n", err)
		os.Exit(1)
	}

	configJSON, err := toJSON(config)
	if err != nil {
		fmt.Printf("設定データのJSON変換エラー: %v\n", err)
		os.Exit(1)
	}

	// テンプレートデータの準備
	data := TemplateData{
		CaseDataJSON:         template.JS(caseDataJSON),
		DocumentsJSON:        template.JS(documentsJSON),
		CaseDocumentsJSON:    template.JS(caseDocumentsJSON),
		DocumentFieldsJSON:   template.JS(documentFieldsJSON),
		FieldsJSON:           template.JS(fieldsJSON),
		ConfigJSON:           template.JS(configJSON),
	}

	// テンプレートの読み込みと実行
	tmpl, err := template.ParseFiles("template.html")
	if err != nil {
		fmt.Printf("テンプレートの読み込みエラー: %v\n", err)
		os.Exit(1)
	}

	// 出力ディレクトリの作成
	outputDir := "output"
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Printf("出力ディレクトリの作成エラー: %v\n", err)
		os.Exit(1)
	}

	// HTMLファイルの生成
	outputPath := filepath.Join(outputDir, "index.html")
	outputFile, err := os.Create(outputPath)
	if err != nil {
		fmt.Printf("出力ファイルの作成エラー: %v\n", err)
		os.Exit(1)
	}
	defer outputFile.Close()

	if err := tmpl.Execute(outputFile, data); err != nil {
		fmt.Printf("テンプレートの実行エラー: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✅ HTMLファイルが生成されました: %s\n", outputPath)
	fmt.Printf("   案件数: %d\n", len(cases))
	fmt.Printf("   書類数: %d\n", len(documents))
	fmt.Printf("   フィールド数: %d\n", len(fields))
}

// JSONファイルを読み込んで案件データをパース
func loadCases(filename string) ([]Case, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var cases []Case
	if err := json.Unmarshal(data, &cases); err != nil {
		return nil, err
	}

	return cases, nil
}

// JSONファイルを読み込んで書類データをパース
func loadDocuments(filename string) ([]Document, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var documents []Document
	if err := json.Unmarshal(data, &documents); err != nil {
		return nil, err
	}

	return documents, nil
}

// JSONファイルを読み込んでフィールドデータをパース
func loadFields(filename string) ([]Field, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var fields []Field
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}

	return fields, nil
}

// JSONファイルを読み込んで案件-書類紐づけデータをパース
func loadCaseDocuments(filename string) ([]CaseDocument, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var caseDocuments []CaseDocument
	if err := json.Unmarshal(data, &caseDocuments); err != nil {
		return nil, err
	}

	return caseDocuments, nil
}

// JSONファイルを読み込んで書類-フィールド紐づけデータをパース
func loadDocumentFields(filename string) ([]DocumentField, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var documentFields []DocumentField
	if err := json.Unmarshal(data, &documentFields); err != nil {
		return nil, err
	}

	return documentFields, nil
}

// JSONファイルを読み込んで設定データをパース
func loadConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

// データをJSON文字列に変換
func toJSON(v interface{}) (string, error) {
	bytes, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

