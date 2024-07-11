package Pages

import (
	"StarNote-editor/Support"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func InitHelpPage(pages *tview.Pages, mainView *tview.Grid) {
	help1 := Support.NewHelpText(Support.HelpText1)
	help2 := Support.NewHelpText(Support.HelpText2)
	help3 := Support.NewHelpText(Support.HelpText3)

	help := tview.NewFrame(help1).SetBorders(1, 1, 0, 0, 2, 2)

	help.SetBorder(true).
		SetTitle("Help").
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyEscape {
				pages.SwitchToPage("main")
				return nil
			} else if event.Key() == tcell.KeyEnter {
				switch {
				case help.GetPrimitive() == help1:
					help.SetPrimitive(help2)
				case help.GetPrimitive() == help2:
					help.SetPrimitive(help3)
				case help.GetPrimitive() == help3:
					help.SetPrimitive(help1)
				}
				return nil
			}
			return event
		})

	pages.AddAndSwitchToPage("main", mainView, true).
		AddPage("help", tview.NewGrid().
			SetColumns(0, 64, 0).
			SetRows(0, 22, 0).
			AddItem(help, 1, 1, 1, 1, 0, 0, true), true, false)
}
