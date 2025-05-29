package main

import (
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	toolWindow := tview.NewTextView().SetText("Select a tool to see the output")
	toolWindow.SetBorder(true).SetTitle("Tool Window")

	// display tools in the tools box
	tools := tview.NewList()
	tools.AddItem("Hello World", "Prints Hello World", '1', func() {
		toolWindow.SetText("Hello World")
	})
	tools.AddItem("Hello World 2", "Prints Hello World 2", '2', func() {
		toolWindow.SetText("Hello World 2")
	})
	tools.SetBorder(true).SetTitle("Tools")

	options := tview.NewBox().SetBorder(true).SetTitle("Options")

	// when I click on a tool, I want to display the output in the tool window
	tools.SetChangedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
		toolWindow.SetText("you selected " + mainText)
	})

	flex := tview.NewFlex()
	flex.AddItem(tools, 0, 1, false)
	flex.AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(options, 0, 1, false).
		AddItem(toolWindow, 0, 3, false), 0, 4, false)

	if err := app.SetRoot(flex, true).SetFocus(tools).Run(); err != nil {
		panic(err)
	}
}
