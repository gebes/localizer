package env

import (
	"github.com/joho/godotenv"
	"os"
	"strings"
)

var (
	SourceLanguage        string
	TargetLanguagesDeepL  []string
	TargetLanguagesGoogle []string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	SourceLanguage = os.Getenv("SOURCE_LANGUAGE")
	t := os.Getenv("TARGET_LANGUAGES_DEEPL")
	if len(t) != 0 {
		TargetLanguagesDeepL = strings.Split(t, ",")
	}
	t = os.Getenv("TARGET_LANGUAGES_GOOGLE")
	if len(t) != 0 {
		TargetLanguagesGoogle = strings.Split(t, ",")
	}
}
