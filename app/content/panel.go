package content

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"gtools/app/vars"
)

func SetContent(a fyne.App, window fyne.Window) {
	cont := container.NewMax()
	title := widget.NewLabel("welcome !")
	intro := widget.NewLabel("this is a mini tools")
	intro.Wrapping = fyne.TextWrapWord

	// bottom
	themes := fyne.NewContainerWithLayout(layout.NewGridLayout(2),
		widget.NewButton("Dark", vars.Dark.Action),
		widget.NewButton("Light", vars.Light.Action),
	)

	navigateBorder := container.NewBorder(nil, themes, nil, nil, BuildTree(a, window, title, intro, cont))
	contentBorder := container.NewVBox(container.NewVBox(title, widget.NewSeparator(), intro), cont)

	split := container.NewHSplit(navigateBorder, contentBorder)
	split.Offset = 0.2
	window.SetContent(split)
}
