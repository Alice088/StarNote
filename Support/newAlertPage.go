package Support

import "github.com/rivo/tview"

// CreateAlertPage Creates "error" page /*
func CreateAlertPage(pages *tview.Pages, text string) {
	errorFrame := tview.NewFrame(
		tview.NewTextView().SetDynamicColors(true).SetText(text)).
		SetBorders(1, 1, 0, 0, 2, 2).
		SetBorder(true).
		SetTitle("Alert!")

	pages.AddPage(
		"error",
		tview.NewGrid().SetColumns(0, 64, 0).SetRows(0, 22, 0).AddItem(errorFrame, 1, 1, 1, 1, 0, 0, true),
		true,
		false)
}
