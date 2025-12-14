# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Japanese document form builder that generates HTML forms from JSON data files. The build process reads JSON data (cases, documents, fields, and their relationships) and embeds them into a single HTML file with embedded JavaScript. The generated HTML is a standalone application that can be distributed to Windows users who don't have Go installed.

## Common Commands

### Build and Run
```bash
# Build the HTML form (requires assets/config.json to exist first)
go run build.go

# View the generated form
open output/index.html
```

### Initial Setup (First Time)
```bash
# Create config file from template
cp assets/config.json.example assets/config.json
# Edit config.json to add API URL and secret key
vim assets/config.json
```

### Windows Distribution Package
```bash
# Build Windows executable
GOOS=windows GOARCH=amd64 go build -o form-builder.exe build.go

# Create release by pushing a version tag
git tag v1.0.0
git push origin v1.0.0
# GitHub Actions will automatically build and create a release
```

## Architecture

### Build Process Flow
1. **build.go** reads multiple JSON files from `assets/`:
   - `case.json`: Cases/scenarios (e.g., "定時株主総会")
   - `documents.json`: Document types (e.g., "議事録")
   - `fields.json`: Form field definitions with metadata (label, type, required, sample, priority)
   - `case_documents.json`: Maps cases to their required documents
   - `document_fields.json`: Maps documents to their required fields
   - `config.json`: API settings and app configuration (gitignored, contains secrets)

2. **Data transformation**: JSON arrays are converted to maps (by ID) for efficient lookups

3. **Template rendering**: `template.html` is processed with Go's `html/template`, embedding all JSON data as JavaScript constants

4. **Output**: Single self-contained `output/index.html` file with all data and logic embedded

### Key Design Patterns

**Field Merging**: When multiple documents share the same field (e.g., "company_name"), the form displays it only once. User inputs are automatically distributed to all documents that need it.

**Priority-based Ordering**: Fields have a `priority` attribute (0-100) that controls display order. Higher priority fields (like "company_name" at 100) appear first.

**Two Selection Modes**:
- **Case mode**: Select a case → documents auto-selected → fields merged
- **Manual mode**: Select documents individually → fields merged

**Data Structure in Output JSON**:
```json
{
  "secret": "api-key",
  "case_id": 1,
  "mode": "case",
  "documents": [
    {
      "document_id": "minutes_regular_meeting",
      "document_name": "定時株主総会議事録",
      "fields": {
        "company_name": "株式会社サンプル",
        "meeting_date": "2024-06-28"
      }
    }
  ]
}
```

### Security Considerations

- `config.json` is gitignored because it contains `secret_key`
- The generated `output/index.html` embeds the secret key in plaintext JavaScript
- This is intended for controlled distribution environments only
- `assets/config.json.example` serves as a template

### File Structure Logic

**JSON Data Files** (`assets/`): Source data that defines all cases, documents, and fields. These are the single source of truth for the form structure.

**Template** (`template.html`): Self-contained HTML with embedded CSS and JavaScript. Contains all form rendering logic, validation, LocalStorage persistence, and API submission code.

**Build Program** (`build.go`): Simple data loader and template processor. No business logic - just reads JSON, converts to maps, and injects into template as JavaScript constants.

**Output** (`output/index.html`): Generated artifact, not tracked in git. Can be distributed as a standalone file.

### Windows Distribution

This project uses GitHub Actions to create Windows release packages automatically when version tags are pushed. The workflow (`.github/workflows/release.yml`) builds `form-builder.exe` and packages it with templates, assets, and documentation into `form-builder-windows.zip`.

## Data Schema Notes

### Field Priority Values
- **100**: Highest priority (e.g., company name, basic identifiers)
- **90-95**: Important metadata (e.g., addresses, dates)
- **50-80**: Standard form fields
- **10**: Default priority

### Field Types
Supported types in `fields.json`: `text`, `date`, `number`, `textarea`, `time`, `select`

### Required vs Optional
The `required` boolean in `fields.json` controls whether a field is mandatory. The HTML template displays required fields with a red asterisk and validates them on submit.

## Common Development Tasks

### Adding New Fields
1. Add field definition to `assets/fields.json` with all metadata
2. Add field ID to relevant documents in `assets/document_fields.json`
3. Rebuild: `go run build.go`

### Adding New Documents
1. Add document to `assets/documents.json` (with Japanese and English names)
2. Map fields to document in `assets/document_fields.json`
3. Link document to cases in `assets/case_documents.json` (if using case mode)
4. Rebuild: `go run build.go`

### Adding New Cases
1. Add case to `assets/case.json` with path array for hierarchy
2. Link case to documents in `assets/case_documents.json`
3. Rebuild: `go run build.go`

### Modifying Form Behavior
Edit `template.html` directly - all form rendering, validation, and submission logic is in the embedded JavaScript (lines 356-874).

### Testing Changes
After running `go run build.go`, open `output/index.html` in a browser. Use browser DevTools to inspect the embedded data constants and debug JavaScript.

## Dependencies

- Go 1.21 or later
- No external Go dependencies (uses only standard library)
- Generated HTML works in modern browsers (requires ES6+ JavaScript support)

## Language Note

This is a Japanese-language application. All user-facing text, labels, and documentation for end users are in Japanese. Code comments and variable names in `build.go` are primarily in Japanese as well.
