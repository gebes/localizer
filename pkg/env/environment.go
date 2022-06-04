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
	TargetLanguagesDeepL = strings.Split(os.Getenv("TARGET_LANGUAGES_DEEPL"), ",")
	TargetLanguagesGoogle = strings.Split(os.Getenv("TARGET_LANGUAGES_GOOGLE"), ",")
}
