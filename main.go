package main

import (
	"log"

	"github.com/marcusolsson/tui-go"
        "github.com/tomoyamachi/multi-translate-shell/translator"
)


var (
	translatee = ""
	labels  = map[string]*tui.Label{}
	languages = []translator.TranslateLanguage{
		{"de", "en"},
		{"fr", "en"},
		{"ja","en"},
	}
)

func createResponseAreaInterface(boxSlice []tui.Widget) []tui.Widget {
	// NewVBoxの引数がinterface型のため
	// interfaceを配列に保存
	// Goの仕様で、interface型の引数はinterfaceそのものを渡さなければならない

	is := make([]tui.Widget, len(boxSlice))
	for i := range boxSlice {
		is[i] = boxSlice[i]
	}
	return is
}

func main() {
	var boxSlice = []tui.Widget{}
	for _, language := range languages {
		// 出力するテキスト
		label := tui.NewLabel("")
		label.SetSizePolicy(tui.Expanding, tui.Expanding)
		labels[language.From + ":" + language.To] =  label
		// テキストを表示する領域
		box := tui.NewVBox(label)
		box.SetTitle(language.From + " <==> " + language.To)
		box.SetBorder(true)
		boxSlice = append(boxSlice, box)
	}

	is := createResponseAreaInterface(boxSlice)
	resp := tui.NewVBox( is... )
	resp.SetSizePolicy(tui.Expanding, tui.Preferred)
	browser := tui.NewHBox(resp)
	browser.SetSizePolicy(tui.Preferred, tui.Expanding)


	inputEntry := tui.NewTextEdit()
	inputEntry.SetText(translatee)
	inputEntry.OnTextChanged(func(e *tui.TextEdit) {
		translatee = e.Text()
	})
	inputBox := tui.NewHBox(inputEntry)
	inputBox.SetTitle("Input English text")
	inputBox.SetBorder(true)

	// focusするbox指定
	tui.DefaultFocusChain.Set(inputEntry)

	root := tui.NewVBox(inputBox, browser)
	theme := tui.NewTheme()
	theme.SetStyle("box.focused.border", tui.Style{Fg: tui.ColorYellow, Bg: tui.ColorDefault})

	ui, err := tui.New(root)
	ui.SetTheme(theme)
	if err != nil {
		log.Fatal(err)
	}
	ui.SetKeybinding("Alt+Enter", func() {
		results := translator.Translate(translatee, languages)

		for _, ts := range results {
			labels[ts.Language].SetText(ts.Result.Export + "\n-----\n" + ts.Result.Reimport)
		}
	})

	ui.SetKeybinding("Esc", func() { ui.Quit() })
	ui.SetKeybinding("Ctrl+c", func() { ui.Quit() })

	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}
}