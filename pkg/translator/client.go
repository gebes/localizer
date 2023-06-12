package translator

import (
	"cloud.google.com/go/translate"
	"context"
	"github.com/DaikiYamakawa/deepl-go"
	"github.com/Gebes/localizer/pkg/env"
	"github.com/Gebes/localizer/pkg/logger"
	"golang.org/x/text/language"
)

var deeplClient *deepl.Client
var googleClient *translate.Client

func init() {
	logger.Debug.Printf("Initializing Google Translate client")
	var err error

	if len(env.TargetLanguagesGoogle) == 0 {
		logger.Debug.Printf("Skipping Google Translate client initialization")
	} else {
		googleClient, err = translate.NewClient(context.Background())
		if err != nil {
			logger.Error.Fatalln("Failed to create Google Translate client:", err)
		}
		logger.Debug.Printf("Google Translate client initialized")
	}

	if len(env.TargetLanguagesDeepL) == 0 {
		logger.Debug.Printf("Skipping DeepL client initialization")

	} else {
		logger.Debug.Printf("Initializing DeepL client")
		deeplClient, err = deepl.New("https://api-free.deepl.com", logger.Debug)

		if err != nil {
			logger.Error.Fatalln("Failed to create DeepL client:", err)
		}
		logger.Debug.Printf("DeepL client initialized")
	}

}

func TranslateGoogle(text []string, sourceLang, targetLang string) ([]translate.Translation, error) {
	sl, err := language.Parse(sourceLang)
	if err != nil {
		return nil, err
	}
	tl, err := language.Parse(targetLang)
	if err != nil {
		return nil, err
	}
	return googleClient.Translate(context.Background(), text, tl, &translate.Options{Source: sl, Model: "nmt"})
}

func TranslateDeepL(sentence, sourceLang, targetLang string) (*deepl.TranslateResponse, error) {
	return deeplClient.TranslateSentence(context.Background(), sentence, sourceLang, targetLang)
}
