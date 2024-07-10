package main

import (
	"StarNote-editor/Support"
	"bufio"
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"os"
	"strings"
)

func main() {
	Support.StarText()
	var filepath string
	var content []byte
	var reader = bufio.NewReader(os.Stdin)

	for {
		fmt.Print("-> ")

		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if text == "exit" {
			os.Exit(0)
		}

		if len(text) != 0 {
			var err error

			filepath = text

			content, err = os.ReadFile(filepath)

			if err != nil {
				fmt.Print("Cannot to read file.")
			}

			break
		}
	}

	app := tview.NewApplication().EnableMouse(true)
	pages := tview.NewPages()

	textArea := tview.NewTextArea().SetPlaceholder("Type something to here...").SetText(string(content), false)
	textArea.SetTitle(filepath).SetBorder(true)

	helpInfo :=
		tview.NewTextView().
			SetText(" Press F1 for help, press Ctrl-C to exit, press Ctrl-S to save")

	textPosition := tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignRight)

	updateInfos := func() {
		fromRow, fromColumn, toRow, toColumn := textArea.GetCursor()
		if fromRow == toRow && fromColumn == toColumn {
			textPosition.SetText(fmt.Sprintf("Row: [yellow]%d[white], Column: [yellow]%d ", fromRow, fromColumn))
		} else {
			textPosition.SetText(fmt.Sprintf("[red]From[white] Row: [yellow]%d[white], Column: [yellow]%d[white] - [red]To[white] Row: [yellow]%d[white], To Column: [yellow]%d ", fromRow, fromColumn, toRow, toColumn))
		}
	}

	textArea.SetMovedFunc(updateInfos)
	updateInfos()

	mainView := tview.NewGrid().
		SetRows(0, 1).
		AddItem(textArea, 0, 0, 1, 2, 0, 0, true).
		AddItem(helpInfo, 1, 0, 1, 1, 0, 0, false).
		AddItem(textPosition, 1, 1, 1, 1, 0, 0, false)

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

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyF1 {
			pages.ShowPage("help") //TODO: Check when clicking outside help window with the mouse. Then clicking help again.
			return nil
		}

		if event.Key() == tcell.KeyCtrlS {
			f, err := os.OpenFile(filepath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
			if err != nil {
				Support.CreateAlertPage(pages, err.Error())
				pages.ShowPage("error")
				return nil
			}

			err = os.Truncate(filepath, 0)
			if err != nil {
				Support.CreateAlertPage(pages, err.Error())
				pages.ShowPage("error")
				return nil
			}

			_, err = f.Write([]byte(textArea.GetText()))
			if err != nil {
				Support.CreateAlertPage(pages, err.Error())
				pages.ShowPage("error")
				return nil
			}

			err = f.Close()
			if err != nil {
				Support.CreateAlertPage(pages, err.Error())
				pages.ShowPage("error")
				return nil
			}

			return nil
		}

		return event
	})

	if err := app.SetRoot(pages, true).Run(); err != nil {
		panic(err)
	}
}
