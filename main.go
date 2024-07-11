package main

import (
	"StarNote-editor/InitModules"
	"StarNote-editor/Pages"
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
	textArea := InitModules.InitTextArea(string(content), filepath)

	helpInfo := tview.NewTextView().SetText("[F1] Help, [Ctrl-S] Save, [Ctrl-/] Settings")
	textPosition := tview.NewTextView().SetDynamicColors(true).SetTextAlign(tview.AlignRight)

	updateInfos := InitModules.InitUpdateInfos(textArea, textPosition)
	textArea.SetMovedFunc(updateInfos)
	updateInfos()

	mainView := InitModules.InitMainView(textArea, helpInfo, textPosition)

	Pages.InitSettingsPage(pages)
	Pages.InitHelpPage(pages, mainView)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyF1 {
			pages.ShowPage("help")
			return nil
		}

		if event.Key() == tcell.KeyF2 {
			pages.ShowPage("settings")
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
