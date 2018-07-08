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

func pararellTranslate(translatee string, languages []TranslateLanguage) []*TranslateResult {
	resultChan := make(chan *TranslateResult)
	for _, language := range languages {
		go setTranslateResult(translatee, language, resultChan)
	}
	var results []*TranslateResult
	valid := true
	for valid{
		results = append(results, <-resultChan)
		valid = (len(results) != len(languages))
	}
	return results
}

func setTranslateResult(translatee string, language TranslateLanguage, resultChan chan *TranslateResult) {
	export := requestTranslate(translatee, language.From)
	reimport := requestTranslate(export, language.To)
	resultChan<-&TranslateResult{language.From, TranslateTexts{export, reimport}}
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
