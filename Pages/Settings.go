package Pages

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func InitSettingsPage(pages *tview.Pages) {
	settingsFrame := tview.NewFrame(
		tview.NewTextView().SetDynamicColors(true).SetText("Se")).
		SetBorders(1, 1, 0, 0, 2, 2).
		SetBorder(true).
		SetTitle("Settings").
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyEscape {
				pages.SwitchToPage("main")
				return nil
			}
			return event
		})

	pages.AddPage(
		"settings",
		tview.NewGrid().SetColumns(0, 64, 0).SetRows(0, 22, 0).AddItem(settingsFrame, 1, 1, 1, 1, 0, 0, true),
		true,
		false)
}
