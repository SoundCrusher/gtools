package content

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func SetContent(a fyne.App, window fyne.Window) {
	cont := container.NewMax()
	title := widget.NewLabel("welcome !")
	intro := widget.NewLabel("this is a mini tools")
	intro.Wrapping = fyne.TextWrapWord

	navigateBorder := container.NewBorder(nil, nil, nil, nil, BuildTree(a, window, title, intro, cont))
	contentBorder := container.NewVBox(container.NewVBox(title, widget.NewSeparator(), intro), cont)

	split := container.NewHSplit(navigateBorder, contentBorder)
	split.Offset = 0.2
	window.SetContent(split)
}