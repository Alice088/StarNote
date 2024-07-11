package Support

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// CreateAlertPage Creates "error" page /*
func CreateAlertPage(pages *tview.Pages, text string) {
	errorFrame := tview.NewFrame(
		tview.NewTextView().SetDynamicColors(true).SetText(text)).
		SetBorders(1, 1, 0, 0, 2, 2).
		SetBorder(true).
		SetTitle("Alert!").
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyEscape {
				pages.SwitchToPage("main")
				pages.RemovePage("error")
				return nil
			}
			return event
		})

	pages.AddPage(
		"error",
		tview.NewGrid().SetColumns(0, 64, 0).SetRows(0, 22, 0).AddItem(errorFrame, 1, 1, 1, 1, 0, 0, true),
		true,
		false)
}
