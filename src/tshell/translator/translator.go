package translator

import (
	"cloud.google.com/go/translate"
	"context"
	"golang.org/x/text/language"
	"log"
)

type TranslateLanguage struct {
	From string
	To   string
}

type TranslateTexts struct {
	Export string
	Reimport string
}

type TranslateResult struct {
	Language string
	Result TranslateTexts
}

func Translate(translatee string, languages []TranslateLanguage) []*TranslateResult {
	return pararellTranslate(translatee, languages)
}

func pararellTranslate(beforeText string, languages []TranslateLanguage) []*TranslateResult {
	var results []*TranslateResult
	for _, language := range languages {
		export := requestTranslate(beforeText, language.From)
		reimport := requestTranslate(export, language.To)
		results = append(results, &TranslateResult{language.From, TranslateTexts{export, reimport}})
	}

	return results
}

func requestTranslate(text string, toLanguage string) string {
	ctx := context.Background()
	client, err := translate.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	target, err := language.Parse(toLanguage)
	if err != nil {
		log.Fatalf("Failed to parse target language: %v", err)
	}
	translations, err := client.Translate(ctx, []string{text}, target, nil)
	if err != nil {
		log.Fatalf("Failed to translate text: %v", err)
	}
	return translations[0].Text;
}
