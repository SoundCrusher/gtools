package feature

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func MakeBoxLayout(_ fyne.Window) fyne.CanvasObject {
	return container.NewVBox()
}
