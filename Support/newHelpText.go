package Support

import "github.com/rivo/tview"

func NewHelpText(text string) *tview.TextView {
	return tview.NewTextView().
		SetDynamicColors(true).
		SetText(text)
}
