package translator

import (
	"cloud.google.com/go/translate"
	"context"
	"fmt"
	"golang.org/x/text/language"
	"log"
)

type TranslateLanguage struct {
	from string
	to   string
}

type TranslateTexts struct {
	export string
	reimport string
}

type TranslateResult struct {
	language string
	result TranslateTexts
}

func Translate() string {
	languages := []TranslateLanguage{
		{"ru", "en"},
		{"ja", "en"},
	}
	text := "You can add some features to this repository."
	return pararellTranslate(text, languages)
}

func pararellTranslate(beforeText string, languages []TranslateLanguage) string {
	var results []TranslateResult

	for _, language := range languages {
		export := requestTranslate(beforeText, language.from)
		reimport := requestTranslate(export, language.to)
		results = append(results, TranslateResult{language.from, TranslateTexts{export, reimport}})
	}

	return "Sample String"
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

	fmt.Printf("Translation: %v\n", translations[0].Text)
	return translations[0].Text;
}
