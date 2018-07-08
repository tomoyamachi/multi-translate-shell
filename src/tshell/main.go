package main

import (
	"log"

	"github.com/marcusolsson/tui-go"
	"tshell/translator"
)

var (
	translatee = ""
	labels  = map[string]*tui.Label{}
	resp *tui.Box
)


func main() {

	languages := []translator.TranslateLanguage{
		{"de", "en"},
		{"fr", "en"},
		{"ja","en"},
	}

	boxArray :=  []tui.Widget{}
	for _, language := range languages {
		// 出力するテキスト
		label := tui.NewLabel("")
		label.SetSizePolicy(tui.Expanding, tui.Expanding)
		labels[language.From] =  label
		// テキストを表示する領域
		box := tui.NewVBox(label)
		box.SetTitle(language.From + " <==> " + language.To)
		box.SetBorder(true)
		boxArray = append(boxArray, box)
	}


	// NewVBoxの引数がinterface型のため
	// interfaceを配列に保存
	// Goの仕様で、interface型の引数はinterfaceそのものを渡さなければならない
	is := make([]tui.Widget, len(boxArray))
	for i := range boxArray {
		is[i] = boxArray[i]
	}

	resp = tui.NewVBox( is... )
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
	ui.SetKeybinding("Enter", func() {
		results := translator.Translate(translatee, languages)

		for _, ts := range results {
			labels[ts.Language].SetText(ts.Result.Export + "\n" + ts.Result.Reimport)
		}
	})

	ui.SetKeybinding("Esc", func() { ui.Quit() })
	ui.SetKeybinding("Ctrl+c", func() { ui.Quit() })

	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}
}
