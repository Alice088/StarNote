package InitModules

import "github.com/rivo/tview"

func InitMainView(primitives ...tview.Primitive) *tview.Grid {
	return tview.NewGrid().
		SetRows(0, 1).
		AddItem(primitives[0], 0, 0, 1, 2, 0, 0, true).
		AddItem(primitives[1], 1, 0, 1, 1, 0, 0, false).
		AddItem(primitives[2], 1, 1, 1, 1, 0, 0, false)
}
