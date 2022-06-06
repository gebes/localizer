<h1 align="center">Localizer</h1>
<p align="center">
<font size="4px">
Command Line Tool for generating <a href="https://www.deepl.com/translator">DeepL</a> and <a href="https://translate.google.com/">Google Translate</a> localizations
</font>
</p>
<p align="center">

<a href="http://golang.org">
    <img src="https://img.shields.io/badge/Made%20with-Go-1f425f.svg" alt="Made with Go">
</a>
<img src="https://img.shields.io/github/go-mod/go-version/Gebes/localizer.svg" alt="Go Version">
<a href="https://pkg.go.dev/github.com/Gebes/localizer">
    <img src="https://img.shields.io/badge/godoc-reference-blue.svg" alt="GoDoc">
</a>
<a href="https://goreportcard.com/report/github.com/Gebes/localizer">
    <img src="https://goreportcard.com/badge/github.com/Gebes/localizer" alt="GoReportCard">
</a>
<a href="https://github.com/Gebes/localizer/blob/master/LICENSE">
    <img src="https://img.shields.io/github/license/Gebes/localizer.svg" alt="License">
</a>
<a href="https://GitHub.com/Gebes/localizer/releases/">
    <img src="https://img.shields.io/github/release/Gebes/localizer" alt="Latest release">
</a>
<a href="https://gocover.io/github.com/Gebes/localizer">
    <img src="https://gocover.io/_badge/github.com/Gebes/localizer" alt="Gocover">
</a>
<a href="https://www.codefactor.io/repository/github/Gebes/localizer">
    <img src="https://www.codefactor.io/repository/github/Gebes/localizer/badge" alt="Gocover">
</a>

<br><br>
<a href="https://gitHub.com/Gebes/localizer/graphs/commit-activity">
<img src="https://img.shields.io/badge/Maintained%3F-yes-green.svg" alt="Maintained?">
</a>
<a href="https://github.com/Gebes">
<img src="https://img.shields.io/badge/Maintainer-Gebes-blue" alt="Maintainer">
</a>
</p>

## ⚠️ Disclaimer
With this tool, you can only generate machine-generated localizations, which may not be 100% accurate. Therefore, always check over the generated localizations.

## 🔥 Get Started

### 💻 Installing the tool
#### 👩‍🔬 Clone the repository
1. Ensure you have the git client installed 
2. Run `git clone github.com/Gebes/localizer`
3. Install the tool with `go install ./cmd/localizer`

#### ⬇️ Download the executable
1. Visit the [releases](https://github.com/Gebes/localizer/releases) page 
2. Download the executable for your operating system and architecture
3. Unarchive with `tar -xf file.tar.gz`

### 🔨 Generate localizations
1. Create a .env file
```dotenv
DEEPL_API_KEY=key:fx
GOOGLE_APPLICATION_CREDENTIALS=./google-cloud.json

SOURCE_LANGUAGE=en
TARGET_LANGUAGES_DEEPL=bg,es,de,ru
TARGET_LANGUAGES_GOOGLE=id,it,ja
```
2. Set `SOURCE_LANGUAGE`, `TARGET_LANGUAGES_DEEPL` and `TARGET_LANGUAGES_GOOGLE` as you desire
   * [List](https://www.deepl.com/de/docs-api/other-functions/listing-supported-languages/) of languages supported by DeepL
   * [List](https://cloud.google.com/translate/docs/languages?hl=de) of languages supported by Google Translate
   * `SOURCE_LANGUAGE` must be a single language and `TARGET_LANGUAGES` can be a list of languages or empty.
3. Provide a
   * [DeepL API Key](https://www.deepl.com/docs-api/accessing-the-api/) if `TARGET_LANGUAGES_DEEPL` is not empty
   * [Google Cloud Json](https://console.cloud.google.com/) with [Cloud Translation API](https://console.cloud.google.com/marketplace/product/google/translate.googleapis.com?q=search&referrer=search&project=place2be-309316) enabled if `TARGET_LANGUAGES_GOOGLE` is not empty
4. Create a `add.yaml` and a file with the `SOURCE_LANGUAGE` (e. g. `en.yaml`)
5. Add text that should be localized to `add.yaml`
```yaml
ui:
  auth:
    button1: Text
    button_sign_in: Sign In
```
6. Generate localizations by running `localizer .` (dot specifies the current folder as target)
   * The tool will take all the values from `add.yaml` and add the values with the same path translated into the corresponding `{lang}.yaml` file
   * After the translations are done `add.yaml` will be cleared
