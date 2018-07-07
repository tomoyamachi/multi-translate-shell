package editor

import (
	"github.com/marcusolsson/tui-go"
	"log"
	"tshell/translator"
)

func Start(defaultString string) {
	buffer := tui.NewTextEdit()
	buffer.SetSizePolicy(tui.Expanding, tui.Expanding)
	buffer.SetText(defaultString)
	buffer.SetFocused(true)
	buffer.SetWordWrap(true)

	status := tui.NewStatusBar("lorem.txt")

	root := tui.NewVBox(buffer, status)

	ui, err := tui.New(root)
	if err != nil {
		log.Fatal(err)
	}

	ui.SetKeybinding("Esc", func() { ui.Quit() })
	ui.SetKeybinding("Ctrl+c", func() { ui.Quit() })
	ui.SetKeybinding("Alt+Enter", func() {requestTranslate(ui, "sampledata")})

	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}
}

func requestTranslate(ui tui.UI, translateString string) {
	ui.Quit()
	translator.Translate()
}