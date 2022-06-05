package localization

import (
	"errors"
	. "localizer/pkg/env"
	"localizer/pkg/file"
	"localizer/pkg/translator"
	"os"
)

func Update(path string) error {
	add, err := ReadFileMap(yamlPath(path, "add"))
	if err != nil {
		return err
	}
	source, err := ReadFileMap(yamlPath(path, SourceLanguage))
	if err != nil {
		return err
	}

	for k, v := range add {
		source[k] = v
	}

	err = writeFileMap(yamlPath(path, SourceLanguage), source)
	if err != nil {
		return err
	}

	for _, language := range TargetLanguagesDeepL {
		target := map[string]string{}
		if _, err = os.Stat(yamlPath(path, language)); err == nil && !errors.Is(err, os.ErrNotExist) {
			target, err = ReadFileMap(yamlPath(path, language))
			if err != nil {
				return err
			}
		}

		for k, v := range add {
			translated, err := translator.TranslateDeepL(v, SourceLanguage, language)
			if err != nil {
				return err
			}
			target[k] = translated.Translations[0].Text
		}

		err = writeFileMap(yamlPath(path, language), target)
		if err != nil {
			return err
		}
	}
	for _, language := range TargetLanguagesGoogle {
		target := map[string]string{}
		if _, err := os.Stat(yamlPath(path, language)); err == nil && !errors.Is(err, os.ErrNotExist) {
			target, err = ReadFileMap(yamlPath(path, language))
			if err != nil {
				return err
			}
		}

		toTranslateKey := make([]string, 0, len(add))
		toTranslateValue := make([]string, 0, len(add))

		for k, v := range add {
			toTranslateKey = append(toTranslateKey, k)
			toTranslateValue = append(toTranslateValue, v)
		}

		translations, err := translator.TranslateGoogle(toTranslateValue, SourceLanguage, language)
		if err != nil {
			return err
		}

		for i, translation := range translations {
			target[toTranslateKey[i]] = translation.Text
		}

		err = writeFileMap(yamlPath(path, language), target)
		if err != nil {
			return err
		}
	}

	err = file.Clear(yamlPath(path, "add"))
	if err != nil {
		return err
	}
	return nil
}
