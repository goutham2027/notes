package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	// Define colors for focused and unfocused states
	focusedBorderColor := tcell.ColorBlue
	focusedTitleColor := tcell.ColorRed
	unfocusedBorderColor := tcell.ColorWhite
	unfocusedTitleColor := tcell.ColorWhite

	toolWindow := tview.NewTextView().SetText("Select a tool to see the output")
	toolWindow.SetBorder(true).SetTitle("Tool Window")
	toolWindow.SetBorderColor(unfocusedBorderColor).SetTitleColor(unfocusedTitleColor)

	// display tools in the tools box
	tools := tview.NewList()
	tools.AddItem("Hello World", "Prints Hello World", '1', func() {
		toolWindow.SetText("Hello World")
	})
	tools.AddItem("Hello World 2", "Prints Hello World 2", '2', func() {
		toolWindow.SetText("Hello World 2")
	})
	tools.SetBorder(true).SetTitle("Tools")
	tools.SetBorderColor(focusedBorderColor).SetTitleColor(focusedTitleColor) // Start focused

	options := tview.NewBox().SetBorder(true).SetTitle("Options")
	options.SetBorderColor(unfocusedBorderColor).SetTitleColor(unfocusedTitleColor)

	// Function to update focus colors
	updateFocusColors := func(focused tview.Primitive) {
		// Reset all to unfocused colors
		tools.SetBorderColor(unfocusedBorderColor).SetTitleColor(unfocusedTitleColor)
		toolWindow.SetBorderColor(unfocusedBorderColor).SetTitleColor(unfocusedTitleColor)
		options.SetBorderColor(unfocusedBorderColor).SetTitleColor(unfocusedTitleColor)

		// Set focused component colors
		switch focused {
		case tools:
			tools.SetBorderColor(focusedBorderColor).SetTitleColor(focusedTitleColor)
		case toolWindow:
			toolWindow.SetBorderColor(focusedBorderColor).SetTitleColor(focusedTitleColor)
		case options:
			options.SetBorderColor(focusedBorderColor).SetTitleColor(focusedTitleColor)
		}
	}

	// when I select a tool by pressing enter, I want to display the output in the tool window
	// change the focus to the tool window
	tools.SetSelectedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
		app.SetFocus(toolWindow)
		updateFocusColors(toolWindow)
		toolWindow.SetText("you selected " + mainText)
	})

	// when escape is pressed, focus on the tools box
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			app.SetFocus(tools)
			updateFocusColors(tools)
		}
		return event
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
