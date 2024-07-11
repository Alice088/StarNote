package InitModules

import "github.com/rivo/tview"

func InitTextArea(content string, title string) *tview.TextArea {
	textArea := tview.NewTextArea().SetPlaceholder("Type something to here...").SetText(content, false)
	textArea.SetTitle(title).SetBorder(true)

	return textArea
}
